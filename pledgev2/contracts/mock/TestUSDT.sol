// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
/// Sepolia demonstration token with EIP-2612 permit support. Not issuer-backed USDT.
contract TestUSDT is ERC20Permit {
    constructor() ERC20("Pledge Test USDT", "USDT") ERC20Permit("Pledge Test USDT") {
        _mint(msg.sender, 31_000_000 * 1e6);
    }
    function decimals() public pure override returns (uint8) { return 6; }
}
