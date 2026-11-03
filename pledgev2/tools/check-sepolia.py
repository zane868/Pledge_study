#!/usr/bin/env python3
"""Read-only verification of the deployed Sepolia system and its broadcast receipts."""
import json
import os
from pathlib import Path
import subprocess
import urllib.request
import http.client
import threading
from urllib.parse import urlsplit

ROOT = Path(__file__).resolve().parent.parent
RPC_URL = os.environ["SEPOLIA_RPC_URL"]


_connections = threading.local()
_rpc_target = urlsplit(RPC_URL)

def rpc(method, params):
    payload = json.dumps({"jsonrpc": "2.0", "id": 1, "method": method, "params": params}).encode()
    target = _rpc_target.path or "/"
    if _rpc_target.query: target += "?" + _rpc_target.query
    for attempt in range(2):
        connection = getattr(_connections, "connection", None)
        if connection is None:
            cls = http.client.HTTPSConnection if _rpc_target.scheme == "https" else http.client.HTTPConnection
            connection = cls(_rpc_target.hostname, _rpc_target.port, timeout=45)
            _connections.connection = connection
        try:
            connection.request("POST", target, payload, {"Content-Type": "application/json"})
            response = connection.getresponse()
            data = response.read()
            if response.status != 200: raise RuntimeError("RPC HTTP status " + str(response.status))
            result = json.loads(data)
            if "error" in result: raise RuntimeError(result["error"].get("message", "RPC error"))
            return result["result"]
        except (http.client.RemoteDisconnected, BrokenPipeError, ConnectionResetError):
            connection.close(); _connections.connection = None
            if attempt: raise


def main():
    manifest = json.loads((ROOT / "deployments/sepolia.json").read_text())
    assert manifest["stableDecimals"] == 6, "USDT must use six decimals"
    broadcast = json.loads((ROOT / "broadcast/DeploySepolia.s.sol/11155111/run-latest.json").read_text())
    assert int(rpc("eth_chainId", []), 16) == manifest["chainId"] == 11155111
    block = rpc("eth_blockNumber", [])
    receipts = broadcast["receipts"]
    assert len(receipts) == len(broadcast["transactions"]) > 0, "Incomplete broadcast"
    from concurrent.futures import ThreadPoolExecutor
    with ThreadPoolExecutor(max_workers=4) as executor:
        live_receipts = list(executor.map(lambda receipt: rpc("eth_getTransactionReceipt", [receipt["transactionHash"]]), receipts))
    for receipt, live in zip(receipts, live_receipts):
        assert live and int(live["status"], 16) == 1, "Failed or missing transaction"
        assert live["blockHash"] == receipt["blockHash"], "Receipt changed"

    artifacts = {
        "multisig": "multiSignature.sol/multiSignature.json",
        "oracle": "BscPledgeOracle.sol/BscPledgeOracle.json",
        "stable": "TestUSDT.sol/TestUSDT.json",
        "weth": "WETH9.sol/WETH9.json",
        "factory": "UniswapV2Factory.sol/UniswapV2Factory.json",
        "router": "UniswapV2Router02.sol/UniswapV2Router02.json",
        "pool": "PledgePool.sol/PledgePool.json",
        "sp": "DebtToken.sol/DebtToken.json",
        "jp": "DebtToken.sol/DebtToken.json",
    }
    for name, artifact_path in artifacts.items():
        artifact = json.loads((ROOT / "out" / artifact_path).read_text())["deployedBytecode"]
        expected = bytearray.fromhex(artifact["object"].removeprefix("0x"))
        actual = bytearray.fromhex(rpc("eth_getCode", [manifest[name], block])[2:])
        assert len(actual) == len(expected) > 0, f"Missing or wrong-sized code: {name}"
        for locations in artifact.get("immutableReferences", {}).values():
            for location in locations:
                start, size = location["start"], location["length"]
                expected[start:start + size] = actual[start:start + size] = bytes(size)
        assert actual == expected, f"Runtime bytecode mismatch: {name}"

    def call(address, signature, *args):
        selector = subprocess.check_output(["cast", "sig", signature], text=True).strip()
        encoded = "".join(f"{int(arg, 16) if isinstance(arg, str) else arg:064x}" for arg in args)
        data = rpc("eth_call", [{"to": address, "data": selector + encoded}, block])[2:]
        return [int(data[i:i + 64], 16) for i in range(0, len(data), 64)]

    def address_equal(value, address):
        assert value == int(address, 16), f"Wrong contract association: {address}"

    pool = manifest["pool"]
    multisig = manifest["multisig"]
    assert call(multisig, "threshold()")[0] == 2
    for index, admin in enumerate(manifest["admins"]):
        address_equal(call(multisig, "signatureOwners(uint256)", index)[0], admin)
    for target in [pool, manifest["oracle"], manifest["sp"], manifest["jp"]]:
        address_equal(call(target, "getMultiSignatureAddress()")[0], multisig)
        application = call(multisig, "getApplicationHash(address,address)", manifest["operator"], target)[0]
        assert call(multisig, "getValidSignature(bytes32,uint256)", application, 0)[0] > 0
    address_equal(call(pool, "oracle()")[0], manifest["oracle"])
    address_equal(call(pool, "swapRouter()")[0], manifest["router"])
    address_equal(call(pool, "feeAddress()")[0], manifest["feeAddress"])
    if manifest["stableDecimals"] == 6:
        assert call(pool, "lendScale(uint256)", 0)[0] == 10**12
        address_equal(call(pool, "lendingAsset()")[0], manifest["stable"])
    assert call(pool, "poolLength()")[0] == 1
    assert call(pool, "getPoolState(uint256)", 0)[0] == 0
    assert call(pool, "minAmount()")[0] == 10**manifest["stableDecimals"]
    assert call(pool, "globalPaused()")[0] == 0
    assert call(pool, "lendFee()")[0] == call(pool, "borrowFee()")[0] == 0
    assert call(manifest["stable"], "decimals()")[0] == manifest["stableDecimals"]
    for name in ["sp", "jp"]:
        assert call(manifest[name], "decimals()")[0] == 18
        assert call(manifest[name], "isMinter(address)", pool)[0] == 1
        assert call(manifest[name], "getMinterLength()")[0] == 1
    info = call(pool, "poolBaseInfo(uint256)", 0)
    assert info[0] == manifest["settleTime"] and info[1] == manifest["endTime"]
    assert info[2] == 20_000_000 and info[6] == 200_000_000
    address_equal(info[7], manifest["stable"])
    assert info[8] == 0
    address_equal(info[10], manifest["sp"])
    address_equal(info[11], manifest["jp"])
    assert call(manifest["oracle"], "getPrice(address)", manifest["stable"])[0] == 10**8
    assert call(manifest["oracle"], "getPrice(address)", "0x0")[0] == 2000 * 10**8
    address_equal(call(manifest["router"], "factory()")[0], manifest["factory"])
    address_equal(call(manifest["router"], "WETH()")[0], manifest["weth"])
    address_equal(call(manifest["factory"], "feeToSetter()")[0], manifest["operator"])
    pair_int = call(manifest["factory"], "getPair(address,address)", manifest["stable"], manifest["weth"])[0]
    assert pair_int != 0
    pair = f"0x{pair_int:040x}"
    pair_artifact = json.loads((ROOT / "out/UniswapV2Factory.sol/UniswapV2Pair.json").read_text())
    assert rpc("eth_getCode", [pair, block]).lower() == pair_artifact["deployedBytecode"]["object"].lower()
    reserves = call(pair, "getReserves()")
    assert sorted(reserves[:2]) == sorted([10**17, 200 * 10**manifest["stableDecimals"]]), "DEX reserves differ from seed liquidity"

    deployment_blocks = [int(receipt["blockNumber"], 16) for receipt in receipts]
    manifest.update({
        "pair": pair,
        "deploymentStartBlock": min(deployment_blocks),
        "deploymentEndBlock": max(deployment_blocks),
        "checkedAtBlock": int(block, 16),
        "successfulTransactions": len(receipts),
        "transactionHashes": [receipt["transactionHash"] for receipt in receipts],
        "onchainChecksPassed": True,
    })
    output = ROOT / "networks/sepolia.json"
    if output.exists():
        previous = json.loads(output.read_text())
        if previous.get("pool", "").lower() == pool.lower():
            for key in ["explorerVerification", "allSourcesVerified", "routerVerificationInputSha256",
                        "runtimeHardfork", "compilerEvmTarget"]:
                if key in previous:
                    manifest[key] = previous[key]
    output.parent.mkdir(exist_ok=True)
    output.write_text(json.dumps(manifest, indent=2) + "\n")
    print(f"Verified {len(receipts)} successful receipts, 10 runtime bytecodes, multisig, permissions, pool and DEX.")
    print(f"Public deployment record: {output}")


if __name__ == "__main__":
    try:
        main()
    except Exception as error:
        message = str(error) or repr(error)
        for name in ["SEPOLIA_RPC_URL", "SEPOLIA_PRIVATE_KEY", "SEPOLIA_PRIVATE_KEY_2", "ETHERSCAN_API_KEY"]:
            value = os.environ.get(name)
            if value:
                message = message.replace(value, "[redacted]")
        raise SystemExit(message)
