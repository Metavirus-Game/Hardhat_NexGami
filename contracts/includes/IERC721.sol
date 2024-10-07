// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @dev ERC721 接口，包含 safeMint 函数
 */
interface IERC721 {
    function safeMint(address to) external;
    function transferFrom(address from, address to, uint256 tokenId) external;
}
