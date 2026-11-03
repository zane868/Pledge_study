// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;
import {Test} from "forge-std/Test.sol";
import {BscPledgeOracle} from "../contracts/pledge/BscPledgeOracle.sol";
import {MockMultiSignature} from "../contracts/mock/MockMultiSignature.sol";
contract FeedStub {
    uint8 public decimals;
    int256 public answer;
    constructor(uint8 precision,int256 value){decimals=precision;answer=value;}
    function latestRoundData() external view returns(uint80,int256,uint256,uint256,uint80){return(1,answer,block.timestamp,block.timestamp,1);}
}
contract OraclePrecisionTest is Test {
    function testQuotePrecisionDoesNotDependOnTokenPrecision() public {
        BscPledgeOracle oracle=new BscPledgeOracle(address(new MockMultiSignature()));
        oracle.setAssetsAggregator(address(1),address(new FeedStub(18,1e18)),6);
        oracle.setAssetsAggregator(address(2),address(new FeedStub(6,2000e6)),18);
        assertEq(oracle.getPrice(address(1)),1e8);
        assertEq(oracle.getPrice(address(2)),2000e8);
    }
    function testBatchPricesUseAssetKeysIncludingNativeEth() public {
        BscPledgeOracle oracle=new BscPledgeOracle(address(new MockMultiSignature()));
        uint256[] memory keys=new uint256[](2);keys[0]=0;keys[1]=uint160(address(123));
        uint256[] memory prices=new uint256[](2);prices[0]=2000e8;prices[1]=1e8;
        oracle.setPrices(keys,prices);
        assertEq(oracle.getPrice(address(0)),2000e8);assertEq(oracle.getPrice(address(123)),1e8);
    }
}
