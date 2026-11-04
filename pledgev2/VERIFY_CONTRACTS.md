# Contract Verification on Sepolia

部署后需在 Etherscan Sepolia 上验证以下合约，支持前端读取 ABI：

```bash
export ETHERSCAN_API_KEY="your_api_key_here"
PROFILE=sepolia

# 1. PledgePool (核心池合约)
forge verify-contract 0x31B807f7D5F8a806d6aB64b0BD8465D56238156e \
  contracts/PledgePool.sol:PledgePool \
  --chain-id 11155111 \
  --watch --etherscan-api-key "$ETHERSCAN_API_KEY"

# 2. multiSignature (多签合约)
forge verify-contract 0x99A4822de5806A6f7f6be559201AeC1641bCBDD0 \
  contracts/multiSignature/multiSignature.sol:multiSignature \
  --chain-id 11155111 \
  --watch --etherscan-api-key "$ETHERSCAN_API_KEY"

# 3. TestUSDT (稳定币)
forge verify-contract 0xF0c87DF88B2905875A23C7afb92Bd7Ba93fCa02a \
  contracts/mock/TestUSDT.sol:TestUSDT \
  --chain-id 11155111 \
  --watch --etherscan-api-key "$ETHERSCAN_API_KEY"

# 4. BscPledgeOracle (预言机)
forge verify-contract 0x2D775e58f80752bd7E13Ef9366E6Bdf904F22B6C \
  contracts/BscPledgeOracle.sol:BscPledgeOracle \
  --chain-id 11155111 \
  --watch --etherscan-api-key "$ETHERSCAN_API_KEY"

# 5. UniswapV2Router02 (DEX 路由)
forge verify-contract 0xB89A09cEE7F3af1C7f873c3E229Be0aa8C736eBF \
  uniswap-v2-periphery-sol6/contracts/UniswapV2Router02.sol:UniswapV2Router02 \
  --chain-id 11155111 \
  --watch --etherscan-api-key "$ETHERSCAN_API_KEY"
```

验证完成后，前端在 `shared/deployment.json` 中的地址即可自动获取 ABI。
