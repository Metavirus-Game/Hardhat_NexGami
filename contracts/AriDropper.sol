// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
abstract contract ERC20Basic {
    uint256 public totalSupply;

    function balanceOf(address who) public view virtual returns (uint256);

    function transfer(address to, uint256 value) public virtual returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
abstract contract ERC20 is ERC20Basic {
    function allowance(
        address owner,
        address spender
    ) public view virtual returns (uint256);

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) public virtual returns (bool);

    function approve(
        address spender,
        uint256 value
    ) public virtual returns (bool);

    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}

contract AirDropper is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {}

    function multisend(
        address _tokenAddr,
        address[] calldata dests,
        uint256[] calldata values
    ) public onlyOwner returns (uint256) {
        uint256 i = 0;
        while (i < dests.length) {
            ERC20(_tokenAddr).transferFrom(msg.sender, dests[i], values[i]);
            i += 1;
        }
        return (i);
    }

	function multisend1(
        address _tokenAddr,
        address[] calldata dests,
        uint256 value
    ) public onlyOwner returns (uint256) {
        uint256 i = 0;
        while (i < dests.length) {
            ERC20(_tokenAddr).transferFrom(msg.sender, dests[i], value);
            i += 1;
        }
        return (i);
    }
}
