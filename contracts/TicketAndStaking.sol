// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TicketAndStaking is Ownable {
    IERC20 public ticketToken; // 门票代币合约
    IERC20 public stakeToken; // 质押代币合约
    IERC20 public rewardToken; // 奖励代币合约

    uint256 public ticketAmount; // 门票代币的数量
    uint256 public minStakeAmount; // 最小质押数量
    uint256 public maxStakeAmount; // 单个地址最大质押数量
    uint256 public totalStakeLimit; // 总质押上限
    uint256 public tgePercentage; // TGE时可领取的百分比
    uint256 public rewardTokenRatio; // 奖励代币比例
    uint256 public startTime; // 开始时间 (Unix 时间戳)
    uint256 public duration; // 持续时间 (秒)
    uint256 public refundDuration; // 退款期限 (秒)

    uint256 public totalStaked; // 当前总质押代币数量

    struct StakeInfo {
        address user;
        uint256 totalStakeAmount;
        bool claimed;
    }

    struct StakingParams {
        IERC20 ticketToken;
        uint256 ticketAmount;
        IERC20 stakeToken;
        uint256 minStakeAmount;
        uint256 maxStakeAmount;
        uint256 totalStakeLimit;
        IERC20 rewardToken;
        uint256 tgePercentage;
        uint256 rewardTokenRatio;
        uint256 startTime;
        uint256 duration;
    }

    mapping(address => StakeInfo) public stakes; // 记录每个用户质押的代币信息
    address[] public participants; // 参与质押的用户地址

    event Staked(address indexed user, uint256 amount);
    event Claimed(
        address indexed user,
        uint256 stakeAmount,
        uint256 rewardAmount
    );

    // constructor(
    //     IERC20 _ticketToken,
    //     uint256 _ticketAmount,
    //     IERC20 _stakeToken,
    //     uint256 _minStakeAmount,
    //     uint256 _maxStakeAmount,
    //     uint256 _totalStakeLimit,
    //     IERC20 _rewardToken,
    //     uint256 _tgePercentage,
    //     uint256 _rewardTokenRatio,
    //     uint256 _startTime,
    //     uint256 _duration,
    //     uint256 _refundDuration
    // ) Ownable(msg.sender) {
    //     ticketToken = _ticketToken;
    //     ticketAmount = _ticketAmount;
    //     stakeToken = _stakeToken;
    //     minStakeAmount = _minStakeAmount;
    //     maxStakeAmount = _maxStakeAmount;
    //     totalStakeLimit = _totalStakeLimit;
    //     rewardToken = _rewardToken;
    //     tgePercentage = _tgePercentage;
    //     rewardTokenRatio = _rewardTokenRatio;
    //     startTime = _startTime;
    //     duration = _duration;
    //     refundDuration = _refundDuration;
    // }

    constructor(StakingParams memory params) Ownable(msg.sender) {
        ticketToken = params.ticketToken;
        ticketAmount = params.ticketAmount;
        stakeToken = params.stakeToken;
        minStakeAmount = params.minStakeAmount;
        maxStakeAmount = params.maxStakeAmount;
        totalStakeLimit = params.totalStakeLimit;
        rewardToken = params.rewardToken;
        tgePercentage = params.tgePercentage;
        rewardTokenRatio = params.rewardTokenRatio;
        startTime = params.startTime;
        duration = params.duration;
    }
    
    function setEventDetails(
        uint256 _startTime,
        uint256 _duration,
        uint256 _refundDuration
    ) external onlyOwner {
        startTime = _startTime;
        duration = _duration;
        refundDuration = _refundDuration;
    }

    function setMaxStakeAmount(uint256 _maxStakeAmount) external onlyOwner {
        maxStakeAmount = _maxStakeAmount;
    }

    function setRewardTokenRatio(uint256 _rewardTokenRatio) external onlyOwner {
        rewardTokenRatio = _rewardTokenRatio;
    }

    function setTgePercentage(uint256 _tgePercentage) external onlyOwner {
        tgePercentage = _tgePercentage;
    }

    function stake(uint256 stakeAmount) external {
        require(block.timestamp >= startTime, "Event has not started yet");
        require(
            block.timestamp <= startTime + duration,
            "Staking period has ended"
        );
        require(
            stakeAmount >= minStakeAmount,
            "Stake amount is less than minimum limit"
        );
        require(
            stakes[msg.sender].totalStakeAmount + stakeAmount <= maxStakeAmount,
            "Stake amount exceeds your maximum stake limit"
        );

        uint256 availableStake = totalStakeLimit - totalStaked;
        require(availableStake > 0, "Total stake limit reached");

        if (stakeAmount > availableStake) {
            stakeAmount = availableStake;
        }

        if (ticketAmount > 0) {
            require(
                ticketToken.allowance(msg.sender, address(this)) >=
                    ticketAmount,
                "Insufficient ticket token allowance"
            );
            require(
                ticketToken.balanceOf(msg.sender) >= ticketAmount,
                "Insufficient ticket token balance"
            );
        }

        require(
            stakeToken.allowance(msg.sender, address(this)) >= stakeAmount,
            "Insufficient stake token allowance"
        );
        require(
            stakeToken.balanceOf(msg.sender) >= stakeAmount,
            "Insufficient stake token balance"
        );

        if (ticketAmount > 0) {
            require(
                ticketToken.transferFrom(
                    msg.sender,
                    address(this),
                    ticketAmount
                ),
                "Ticket token transfer failed"
            );
        }
        require(
            stakeToken.transferFrom(msg.sender, address(this), stakeAmount),
            "Stake token transfer failed"
        );

        if (stakes[msg.sender].user == address(0)) {
            stakes[msg.sender].user = msg.sender;
            participants.push(msg.sender);
        }

        stakes[msg.sender].totalStakeAmount += stakeAmount;
        totalStaked += stakeAmount;

        emit Staked(msg.sender, stakeAmount);
    }

    function claim() external {
        require(block.timestamp > startTime + duration, "Event has not ended");
        require(
            block.timestamp <= startTime + duration + refundDuration,
            "Claim period has ended"
        );
        uint256 stakeAmount = stakes[msg.sender].totalStakeAmount;
        require(stakeAmount > 0, "No staked amount to claim");

        uint256 rewardAmount = (stakeAmount * tgePercentage) / 100;
        uint256 rewardTokenAmount = (stakeAmount * rewardTokenRatio) / 100;

        require(
            stakeToken.transfer(msg.sender, stakeAmount),
            "Stake token transfer failed"
        );
        require(
            rewardToken.transfer(msg.sender, rewardTokenAmount),
            "Reward token transfer failed"
        );

        stakes[msg.sender].claimed = true;

        emit Claimed(msg.sender, stakeAmount, rewardAmount);
    }

    function getStakingInfo()
        external
        view
        returns (uint256, uint256, uint256, bool)
    {
        StakeInfo memory stakeInfo = stakes[msg.sender];
        return (
            totalStaked,
            participants.length,
            stakeInfo.totalStakeAmount,
            stakeInfo.claimed
        );
    }

    function withdrawTicketTokens(uint256 amount) external onlyOwner {
        require(
            ticketToken.transfer(owner(), amount),
            "Ticket token transfer failed"
        );
    }

    function withdrawRewardTokens(uint256 amount) external onlyOwner {
        require(
            rewardToken.transfer(owner(), amount),
            "Reward token transfer failed"
        );
    }

    function withdrawStakeTokens(uint256 amount) external onlyOwner {
        require(
            stakeToken.transfer(owner(), amount),
            "Stake token transfer failed"
        );
    }

    function getContractBalances()
        external
        view
        onlyOwner
        returns (
            uint256 ticketBalance,
            uint256 rewardBalance,
            uint256 stakeBalance
        )
    {
        ticketBalance = ticketToken.balanceOf(address(this));
        rewardBalance = rewardToken.balanceOf(address(this));
        stakeBalance = stakeToken.balanceOf(address(this));
    }

    function getClaimedUsers()
        external
        view
        onlyOwner
        returns (address[] memory)
    {
        address[] memory claimedUsers = new address[](participants.length);
        uint256 count = 0;
        for (uint256 i = 0; i < participants.length; i++) {
            if (stakes[participants[i]].claimed) {
                claimedUsers[count] = participants[i];
                count++;
            }
        }
        // Resize the array to the actual number of claimed users
        assembly {
            mstore(claimedUsers, count)
        }
        return claimedUsers;
    }
}

// contract TicketAndStaking is Ownable {
//     IERC20 public ticketToken; // 门票代币合约
//     IERC20 public stakeToken; // 质押代币合约
//     IERC20 public rewardToken; // 奖励代币合约

//     uint256 public ticketAmount; // 门票代币的数量
//     uint256 public minStakeAmount; // 最小质押数量
//     uint256 public maxStakeAmount; // 单个地址最大质押数量
//     uint256 public totalStakeLimit; // 总质押上限
//     uint256 public rewardRatio; // 奖励比例
//     uint256 public tgePercentage; //tge发放额度，剩余部分vesting
//     uint256 public startTime; // 开始时间 (Unix 时间戳)
//     uint256 public duration; // 持续时间 (秒)
//     uint256 public refundDuration; // 退款期限 (秒)

//     uint256 public totalStaked; // 当前总质押代币数量
//     uint256 public participantCount; // 参与人数统计

//     mapping(address => uint256) public stakes; // 记录每个用户质押的代币数量

//     event Staked(address indexed user, uint256 amount);
//     event Claimed(
//         address indexed user,
//         uint256 stakeAmount,
//         uint256 rewardAmount
//     );
//     event Refunded(address indexed user, uint256 stakeAmount);

//     constructor(
//         IERC20 _ticketToken,
//         uint256 _ticketAmount,
//         IERC20 _stakeToken,
//         uint256 _minStakeAmount,
//         uint256 _maxStakeAmount,
//         uint256 _totalStakeLimit,
//         IERC20 _rewardToken,
//         uint256 _rewardRatio,
//         uint256 _tgePercentage,
//         uint256 _startTime,
//         uint256 _duration,
//         uint256 _refundDuration
//     ) Ownable(msg.sender) {
//         ticketToken = _ticketToken;
//         ticketAmount = _ticketAmount;
//         stakeToken = _stakeToken;
//         minStakeAmount = _minStakeAmount;
//         maxStakeAmount = _maxStakeAmount;
//         totalStakeLimit = _totalStakeLimit;
//         rewardToken = _rewardToken;
//         tgePercentage = _tgePercentage;
//         rewardRatio = _rewardRatio;
//         startTime = _startTime;
//         duration = _duration;
//         refundDuration = _refundDuration;
//     }

//     function setMaxStakeAmount(uint256 _maxStakeAmount) external onlyOwner {
//         maxStakeAmount = _maxStakeAmount;
//     }

//     function stake(uint256 stakeAmount) external {
//         require(block.timestamp >= startTime, "Event has not started yet");
//         require(
//             block.timestamp <= startTime + duration,
//             "Staking period has ended"
//         );
//         require(
//             stakeAmount >= minStakeAmount,
//             "Stake amount is less than minimum limit"
//         );
//         require(
//             stakes[msg.sender] + stakeAmount <= maxStakeAmount,
//             "Stake amount exceeds your maximum stake limit"
//         );

//         uint256 availableStake = totalStakeLimit - totalStaked;
//         require(availableStake > 0, "Total stake limit reached");

//         if (stakeAmount > availableStake) {
//             stakeAmount = availableStake;
//         }
//         if (ticketAmount > 0) {
//             require(
//                 ticketToken.allowance(msg.sender, address(this)) >=
//                     ticketAmount,
//                 "Insufficient ticket token allowance"
//             );
//         }
//         require(
//             stakeToken.allowance(msg.sender, address(this)) >= stakeAmount,
//             "Insufficient stake token allowance"
//         );

//         if (ticketAmount > 0) {
//             ticketToken.transferFrom(msg.sender, address(this), ticketAmount);
//         }
//         stakeToken.transferFrom(msg.sender, address(this), stakeAmount);

//         if (stakes[msg.sender] == 0) {
//             participantCount++;
//         }

//         stakes[msg.sender] += stakeAmount;
//         totalStaked += stakeAmount;

//         emit Staked(msg.sender, stakeAmount);
//     }

//     function claim() external {
//         require(block.timestamp > startTime + duration, "Event has not ended");
//         require(
//             block.timestamp <= startTime + duration + refundDuration,
//             "Claim period has ended"
//         );
//         uint256 stakeAmount = stakes[msg.sender];
//         require(stakeAmount > 0, "No staked amount to claim");

//         uint256 rewardAmount = (stakeAmount * rewardRatio) / 100;
//         stakes[msg.sender] = 0;
//         //totalStaked -= stakeAmount;

//         require(
//             stakeToken.transfer(msg.sender, stakeAmount),
//             "Stake token transfer failed"
//         );
//         require(
//             rewardToken.transfer(msg.sender, rewardAmount),
//             "Reward token transfer failed"
//         );

//         emit Claimed(msg.sender, stakeAmount, rewardAmount);
//     }

//     function refund() external {
//         require(block.timestamp > startTime + duration, "Event has not ended");
//         require(
//             block.timestamp <= startTime + duration + refundDuration,
//             "Refund period has ended"
//         );
//         uint256 stakeAmount = stakes[msg.sender];
//         require(stakeAmount > 0, "No staked amount to refund");

//         stakes[msg.sender] = 0;
//         //totalStaked -= stakeAmount;

//         require(
//             stakeToken.transfer(msg.sender, stakeAmount),
//             "Stake token transfer failed"
//         );

//         emit Refunded(msg.sender, stakeAmount);
//     }

//     function getStakingInfo()
//         external
//         view
//         returns (uint256, uint256, uint256)
//     {
//         return (totalStaked, participantCount, stakes[msg.sender]);
//     }

//     function getStakeAmount(address user) external view returns (uint256) {
//         return stakes[user];
//     }

//     function setRewardRatio(uint256 _rewardRatio) external onlyOwner {
//         rewardRatio = _rewardRatio;
//     }

//     function withdrawTicketTokens(uint256 amount) external onlyOwner {
//         require(
//             ticketToken.transfer(owner(), amount),
//             "Ticket token transfer failed"
//         );
//     }

//     function withdrawRewardTokens(uint256 amount) external onlyOwner {
//         require(
//             rewardToken.transfer(owner(), amount),
//             "Reward token transfer failed"
//         );
//     }

//     function withdrawStakeTokens(uint256 amount) external onlyOwner {
//         require(
//             stakeToken.transfer(owner(), amount),
//             "Stake token transfer failed"
//         );
//     }

//     function getContractBalances()
//         external
//         view
//         onlyOwner
//         returns (
//             uint256 ticketBalance,
//             uint256 rewardBalance,
//             uint256 stakeBalance
//         )
//     {
//         ticketBalance = ticketToken.balanceOf(address(this));
//         rewardBalance = rewardToken.balanceOf(address(this));
//         stakeBalance = stakeToken.balanceOf(address(this));
//     }
// }
