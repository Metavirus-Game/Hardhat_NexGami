// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @dev 扩展的 IERC20 接口，包含 decimals 函数
 */
interface IERC20Extended is IERC20 {
    function decimals() external view returns (uint8);
}
