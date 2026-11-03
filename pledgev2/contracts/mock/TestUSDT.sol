// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
/// Sepolia demonstration token, not issuer-backed USDT.
contract TestUSDT is ERC20 {
    constructor() ERC20("Pledge Test USDT", "USDT") { _mint(msg.sender, 31_000_000 * 1e6); }
    function decimals() public pure override returns (uint8) { return 6; }
}
