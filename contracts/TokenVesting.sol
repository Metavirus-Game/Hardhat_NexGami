// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract TokenVesting is Ownable, ReentrancyGuard {
    event TokensReleased(address token, uint256 amount);
    event TokenVestingRevoked(address token);

    IERC20 private _token;
    address private _beneficiary;
    uint256 private _start;
    uint256 private _cliff;
    uint256 private _duration;
    uint256 private _released;
    bool private _revoked;

    constructor(
        IERC20 token_,
        address beneficiary_,
        uint256 start_,
        uint256 cliffDuration_,
        uint256 duration_
    ) Ownable(msg.sender) {
        require(
            address(token_) != address(0),
            "TokenVesting: token is the zero address"
        );
        require(
            beneficiary_ != address(0),
            "TokenVesting: beneficiary is the zero address"
        );
        require(duration_ > 0, "TokenVesting: duration is 0");
        require(
            cliffDuration_ <= duration_,
            "TokenVesting: cliff is longer than duration"
        );

        _token = token_;
        _beneficiary = beneficiary_;
        _start = start_;
        _cliff = start_ + cliffDuration_;
        _duration = duration_;
        _released = 0;
        _revoked = false;
    }

    function release() public nonReentrant {
        require(!_revoked, "TokenVesting: contract is revoked");
        uint256 unreleased = _releasableAmount();
        require(unreleased > 0, "TokenVesting: no tokens are due");

        _released += unreleased;

        _token.transfer(_beneficiary, unreleased);

        emit TokensReleased(address(_token), unreleased);
    }

    function revoke() public onlyOwner {
        require(!_revoked, "TokenVesting: contract already revoked");
        uint256 balance = _token.balanceOf(address(this));
        uint256 unreleased = _releasableAmount();
        uint256 refund = balance - unreleased;

        _revoked = true;
        _token.transfer(owner(), refund);

        emit TokenVestingRevoked(address(_token));
    }

    function releasableAmount() public view returns (uint256) {
        return _releasableAmount();
    }

    function _releasableAmount() private view returns (uint256) {
        if (_revoked) {
            return 0;
        }
        return _vestedAmount() - _released;
    }

    function _vestedAmount() private view returns (uint256) {
        if (block.timestamp < _cliff) {
            return 0;
        } else if (block.timestamp >= _start + _duration) {
            return _token.balanceOf(address(this));
        } else {
            return
                (_token.balanceOf(address(this)) * (block.timestamp - _start)) /
                _duration;
        }
    }
}
