// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Vm} from "forge-std/Vm.sol";

interface IToken {
    function approve(address, uint256) external returns (bool);
    function transfer(address, uint256) external returns (bool);
    function balanceOf(address) external view returns (uint256);
    function totalSupply() external view returns (uint256);
    function addMinter(address) external returns (bool);
    function delMinter(address) external returns (bool);
    function isMinter(address) external view returns (bool);
    function mint(address, uint256) external returns (bool);
    function burn(address, uint256) external returns (bool);
}

interface IMultisig {
    function createApplication(address) external returns (uint256);
    function getApplicationHash(address, address) external pure returns (bytes32);
    function signApplication(bytes32) external;
    function revokeSignApplication(bytes32) external;
}

interface IOracle {
    function setPrice(address, uint256) external;
}

interface IRouter {
    function addLiquidityETH(address, uint256, uint256, uint256, address, uint256)
        external
        payable
        returns (uint256, uint256, uint256);
    function getAmountsIn(uint256, address[] calldata) external view returns (uint256[] memory);
    function getAmountsOut(uint256, address[] calldata) external view returns (uint256[] memory);
}

interface IPool {
    function createPoolInfo(
        uint256,
        uint256,
        uint64,
        uint256,
        uint256,
        address,
        address,
        address,
        address,
        uint256
    ) external;
    function poolLength() external view returns (uint256);
    function poolDataInfo(uint256)
        external
        view
        returns (uint256, uint256, uint256, uint256, uint256, uint256);
    function userBorrowInfo(address, uint256) external view returns (uint256, uint256, bool, bool);
    function getMultiSignatureAddress() external view returns (address);
    function oracle() external view returns (address);
    function borrowFee() external view returns (uint256);
    function getPoolState(uint256) external view returns (uint256);
    function checkoutSettle(uint256) external view returns (bool);
    function checkoutLiquidate(uint256) external view returns (bool);
    function depositLend(uint256, uint256) external payable;
    function depositBorrow(uint256, uint256) external payable;
    function refundLend(uint256) external;
    function refundBorrow(uint256) external;
    function claimLend(uint256) external;
    function claimBorrow(uint256) external;
    function withdrawLend(uint256, uint256) external;
    function withdrawBorrow(uint256, uint256) external;
    function emergencyLendWithdrawal(uint256) external;
    function emergencyBorrowWithdrawal(uint256) external;
    function setFee(uint256, uint256) external;
    function setMinAmount(uint256) external;
    function setSwapRouterAddress(address) external;
    function setPause() external;
    function settle(uint256) external;
    function finish(uint256) external;
    function liquidate(uint256) external;
}

/// Shared setup for Forge tests and the Anvil deployment script.
/// Explicit artifact paths distinguish the formal pool from similarly named mock artifacts.
abstract contract LocalSystem {
    Vm internal constant cheat = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));
    uint256 internal constant YEAR = 365 days;
    address internal operator;
    address internal ownerA;
    address internal ownerB;
    address internal lender;
    address internal borrower;
    address internal recipient;
    address internal fee;
    IMultisig internal multisig;
    IOracle internal oracle;
    IToken internal stable;
    IToken internal sp;
    IToken internal jp;
    IPool internal pool;
    IRouter internal router;
    address internal weth;
    address internal factory;
    uint256 internal pid;
    uint256 internal settleTime;
    uint256 internal endTime;

    function _start(address sender) internal virtual;
    function _stop() internal virtual;

    function _liquidityEth() internal view virtual returns (uint256) {
        return 1000 ether;
    }

    function _poolDuration() internal view virtual returns (uint256) {
        return YEAR;
    }

    function _poolDelay() internal view virtual returns (uint256) {
        return 1 hours;
    }

    function _stableUnit() internal pure virtual returns (uint256) { return 1 ether; }

    function _stableArtifact() internal pure virtual returns (string memory) { return "out/BUSD.sol/BEP20Token.json"; }

    function _signatureOwners() internal view virtual returns (address[] memory owners) {
        owners = new address[](2);
        owners[0] = ownerA;
        owners[1] = ownerB;
    }

    function _deploy(string memory artifact, bytes memory args) internal returns (address deployed) {
        bytes memory code = abi.encodePacked(cheat.getCode(artifact), args);
        assembly { deployed := create(0, add(code, 32), mload(code)) }
        require(deployed != address(0), "Local deployment failed");
    }

    function _authorize(address applicant, address target) internal returns (bytes32 hash) {
        _start(applicant);
        multisig.createApplication(target);
        _stop();
        hash = multisig.getApplicationHash(applicant, target);
        _start(ownerA);
        multisig.signApplication(hash);
        _stop();
        _start(ownerB);
        multisig.signApplication(hash);
        _stop();
    }

    function _deploySystem() internal {
        address[] memory owners = _signatureOwners();
        _start(operator);
        multisig = IMultisig(_deploy("out/multiSignature.sol/multiSignature.json", abi.encode(owners, 2)));
        oracle =
            IOracle(_deploy("out/BscPledgeOracle.sol/BscPledgeOracle.json", abi.encode(address(multisig))));
        stable = IToken(_deploy(_stableArtifact(), ""));
        weth = _deploy("out/WETH9.sol/WETH9.json", "");
        factory = _deploy("out/UniswapV2Factory.sol/UniswapV2Factory.json", abi.encode(operator));
        router =
            IRouter(_deploy("out/UniswapV2Router02.sol/UniswapV2Router02.json", abi.encode(factory, weth)));
        pool = IPool(
            _deploy(
                "out/PledgePool.sol/PledgePool.json",
                abi.encode(address(oracle), address(router), fee, address(multisig))
            )
        );
        _stop();
        _authorize(operator, address(oracle));
        _authorize(operator, address(pool));
        _start(operator);
        pool.setMinAmount(100 * _stableUnit());
        _stop();
        _start(operator);
        oracle.setPrice(address(stable), 1e8);
        oracle.setPrice(address(0), 2000e8);
        uint256 liquidityEth = _liquidityEth();
        uint256 liquidityStable = liquidityEth * 2000 * _stableUnit() / 1 ether;
        stable.approve(address(router), liquidityStable);
        router.addLiquidityETH{value: liquidityEth}(
            address(stable), liquidityStable, 0, 0, operator, block.timestamp + 1 hours
        );
        stable.transfer(lender, 20_000 * _stableUnit());
        _stop();
        _createPool(_poolDuration());
    }

    function _createPool(uint256 duration) internal {
        pid = pool.poolLength();
        _start(operator);
        sp = IToken(
            _deploy("out/DebtToken.sol/DebtToken.json", abi.encode("Pledge SP", "SP", address(multisig)))
        );
        jp = IToken(
            _deploy("out/DebtToken.sol/DebtToken.json", abi.encode("Pledge JP", "JP", address(multisig)))
        );
        _stop();
        _authorize(operator, address(sp));
        _authorize(operator, address(jp));
        settleTime = block.timestamp + _poolDelay();
        endTime = settleTime + duration;
        _start(operator);
        sp.addMinter(address(pool));
        jp.addMinter(address(pool));
        pool.createPoolInfo(
            settleTime,
            endTime,
            20_000_000,
            100_000 * _stableUnit(),
            200_000_000,
            address(stable),
            address(0),
            address(sp),
            address(jp),
            20_000_000
        );
        _stop();
    }
}
