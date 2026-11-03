.PHONY: build test fmt node deploy-local deploy-sepolia check-sepolia abi

build:
	forge build

test:
	forge test -vv

fmt:
	forge fmt test script

node:
	anvil

deploy-local:
	forge script script/DeployLocal.s.sol:DeployLocal --rpc-url http://127.0.0.1:8545 --broadcast

deploy-sepolia:
	FOUNDRY_PROFILE=sepolia forge script script/DeploySepolia.s.sol:DeploySepolia --rpc-url "$$SEPOLIA_RPC_URL" --broadcast --slow --build-info --verify --non-interactive

check-sepolia:
	python3 tools/check-sepolia.py

abi: build
	./tools/export-abi
