#!/usr/bin/env python3
"""Regenerate bindings from this repository's Foundry build (abigen 1.17.8)."""
import json
from pathlib import Path
import subprocess
root = Path(__file__).resolve().parent.parent
contracts = [
 ("PledgePool.sol/PledgePool.json", "pledge_pool.abi", "PledgePoolToken", "pledgePoolToken.go"),
 ("BscPledgeOracle.sol/BscPledgeOracle.json", "bsc_pledge_oracle_testnet.abi", "BscPledgeOracleTestnetToken", "bscPledgeOracleTestNetToken.go"),
 ("BscPledgeOracle.sol/BscPledgeOracle.json", "bsc_pledge_oracle_mainnet.abi", "BscPledgeOracleMainnetToken", "bscPledgeOracleMainNetToken.go"),
 ("DebtToken.sol/DebtToken.json", "debt_token.abi", "DebtToken", "debtToken.go"),
 ("multiSignature.sol/multiSignature.json", "multi_signature.abi", "MultiSignature", "multiSignature.go"),
]
for artifact, filename, typename, output in contracts:
    abi = json.loads((root.parent / "pledgev2/out" / artifact).read_text())["abi"]
    path = root / "contract/abi" / filename
    path.write_text(json.dumps(abi, indent=2) + "\n")
    subprocess.run(["abigen", "--abi", str(path), "--pkg", "bindings", "--type", typename,
                    "--out", str(root / "contract/bindings" / output)], check=True)
manifest = json.loads((root.parent / "pledgev2/networks/sepolia.json").read_text())
keys = ["stableDecimals", "chainId", "pool", "oracle", "multisig", "stable", "weth", "router", "factory", "sp", "jp", "operator", "admins", "threshold", "deploymentStartBlock"]
(root / "config/sepolia.json").write_text(json.dumps({key: manifest[key] for key in keys}, indent=2) + "\n")
