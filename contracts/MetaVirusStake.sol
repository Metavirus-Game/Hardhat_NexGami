// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./includes/IERC721.sol";

interface IERC20Permit is IERC20 {
    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;
}

contract MetaVirusVIPStake is Ownable {
    IERC20Permit public stakingToken;
    uint256 public rewardRate; // 奖励率（基于万分比的值）
    uint256 public stakingDuration; // 质押时长（以秒为单位）
    IERC721 public rewardNFT; // 奖励NFT

    struct Stake {
        uint256 amount; // 质押总数量
        uint256 startTime; // 最后一次质押开始时间
        bool claimed; // 是否已经领取奖励
    }

    mapping(address => Stake) public stakes;

    event Staked(address indexed user, uint256 amount, uint256 startTime);
    event Claimed(address indexed user, uint256 amount, uint256 reward);
    event NFTRewarded(address indexed user, uint256 tokenId); // NFT奖励事件

    constructor(
        IERC20Permit _stakingToken,
        uint256 _rewardRate,
        address _rewardNFT,
        uint256 _stakingDuration
    ) Ownable(msg.sender) {
        stakingToken = _stakingToken;
        rewardRate = _rewardRate;
        rewardNFT = IERC721(_rewardNFT);
        stakingDuration = _stakingDuration;
    }

    function stake(uint256 amount) external {
        require(amount > 0, "Staking amount must be greater than zero");

        stakingToken.transferFrom(msg.sender, address(this), amount);

        if (stakes[msg.sender].amount == 0) {
            //发放NFT
            rewardNFT.safeMint(msg.sender);
            emit NFTRewarded(msg.sender, 0);
        }

        stakes[msg.sender].amount += amount;
        stakes[msg.sender].startTime = block.timestamp;
        stakes[msg.sender].claimed = false;

        emit Staked(msg.sender, amount, block.timestamp);
    }

    function stakeWithPermit(
        address user,
        uint256 amount,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external {
        require(amount > 0, "Staking amount must be greater than zero");

        stakingToken.permit(user, address(this), amount, deadline, v, r, s);
        stakingToken.transferFrom(user, address(this), amount);

        stakes[user].amount += amount;
        stakes[user].startTime = block.timestamp;
        stakes[user].claimed = false;

        emit Staked(user, amount, block.timestamp);
    }

    function claim(address user) external {
        Stake storage userStake = stakes[user];
        require(userStake.amount > 0, "No staked amount");
        require(!userStake.claimed, "Already claimed");
        require(
            block.timestamp >= userStake.startTime + stakingDuration,
            "Staking period not yet finished"
        );

        uint256 reward = (userStake.amount * rewardRate) / 10000;
        uint256 totalAmount = userStake.amount + reward;
        userStake.claimed = true;

        stakingToken.transfer(user, totalAmount);

        emit Claimed(user, userStake.amount, reward);
    }

    function withdraw(
        uint256 tokenAmount,
        uint256 maticAmount
    ) external onlyOwner {
        require(
            tokenAmount <= stakingToken.balanceOf(address(this)),
            "Insufficient token balance"
        );
        require(
            maticAmount <= address(this).balance,
            "Insufficient MATIC balance"
        );

        stakingToken.transfer(owner(), tokenAmount);
        payable(owner()).transfer(maticAmount);
    }

    function updateStakingToken(IERC20Permit newToken) external onlyOwner {
        stakingToken = newToken;
    }

    function updateRewardRate(uint256 newRate) external onlyOwner {
        rewardRate = newRate;
    }

    function updateStakingDuration(uint256 newDuration) external onlyOwner {
        stakingDuration = newDuration;
    }

    receive() external payable {}
}
