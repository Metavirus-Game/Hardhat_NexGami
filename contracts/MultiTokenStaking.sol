// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./includes/IERC20Extended.sol";
import "./includes/IERC721.sol";

/**
 * @title MultiTokenStaking
 * @dev 允许用户质押基础代币和其他代币，根据年化利率产出指定代币，并在满足条件时奖励NFT。
 */
contract MultiTokenStaking is Ownable, ReentrancyGuard {
    struct StakeInfo {
        uint256 startTime; // 质押开始时间
        uint256 lastClaimTime; // 上次领取奖励的时间
        mapping(IERC20 => uint256) tokenStaked; // 用户质押的代币数量，包括基础代币和其他代币
        bool nftClaimed; // 用户是否已经领取过NFT奖励
    }

    IERC20Extended public baseToken; // 基础代币
    IERC20 public rewardToken; // 奖励代币
    IERC721 public rewardNFT; // 奖励NFT
    uint256 public annualInterestRate; // 年化利率，单位为百分之一，例如年化利率5%则为500
    uint256 public nftThreshold; // 获得NFT的最小质押数量
    uint256 public minStakeDuration; // 获得奖励的最短质押时长
    uint256 public minStakeAmount; // 最小质押数量

    uint256 public totalBaseStaked; // 基础代币的质押总量

    mapping(address => StakeInfo) public stakes; // 用户地址到质押信息的映射
    mapping(IERC20 => uint256) public otherTokenRatios; // 其他代币与基础代币的比例，单位为wei
    mapping(IERC20 => uint8) public tokenDecimals; // 记录代币的精度
    IERC20[] public otherTokens; // 记录所有的其他代币

    event Staked(
        address indexed user,
        uint256 amount,
        uint256 startTime,
        uint256 totalBaseStaked,
        uint8 nftStatus
    ); // 质押事件
    event Unstaked(
        address indexed user,
        uint256 amount,
        uint256 totalBaseStaked,
        uint8 nftStatus
    ); // 解除质押事件
    event RewardClaimed(
        address indexed user,
        uint256 rewardAmount,
        uint256 timestamp
    ); // 奖励领取事件
    event NFTRewarded(address indexed user, uint256 tokenId); // NFT奖励事件

    /**
     * @dev 构造函数，初始化合约
     * @param _baseToken 基础代币地址
     * @param _rewardToken 奖励代币地址
     * @param _annualInterestRate 年化利率
     * @param _nftThreshold 获得NFT的最小质押数量
     * @param _minStakeDuration 获得奖励的最短质押时长
     * @param _minStakeAmount 最小质押数量
     */
    constructor(
        address _baseToken,
        address _rewardToken,
        uint256 _annualInterestRate,
        uint256 _nftThreshold,
        uint256 _minStakeDuration,
        uint256 _minStakeAmount
    ) Ownable(msg.sender) {
        baseToken = IERC20Extended(_baseToken);
        rewardToken = IERC20(_rewardToken);
        annualInterestRate = _annualInterestRate;
        nftThreshold = _nftThreshold;
        minStakeDuration = _minStakeDuration;
        minStakeAmount = _minStakeAmount;

        try baseToken.decimals() returns (uint8 decimals) {
            tokenDecimals[baseToken] = decimals;
        } catch {
            tokenDecimals[baseToken] = 18; // 默认18位小数
        }
    }

    /**
     * @dev 配置质押代币和比例，以及奖励NFT地址
     * @param _otherTokens 其他代币数组
     * @param _ratios 其他代币与基础代币的比例数组，单位为wei
     * @param _rewardNFT 奖励NFT地址
     */
    function configureStaking(
        IERC20[] calldata _otherTokens,
        uint256[] calldata _ratios,
        address _rewardNFT
    ) external onlyOwner {
        require(
            _otherTokens.length == _ratios.length,
            "Token and ratio length mismatch"
        );

        rewardNFT = IERC721(_rewardNFT);

        for (uint256 i = 0; i < _otherTokens.length; i++) {
            IERC20Extended otherToken = IERC20Extended(
                address(_otherTokens[i])
            );

            if (otherTokenRatios[_otherTokens[i]] == 0) {
                otherTokens.push(_otherTokens[i]);
            }

            otherTokenRatios[_otherTokens[i]] = _ratios[i];

            try otherToken.decimals() returns (uint8 decimals) {
                tokenDecimals[otherToken] = decimals;
            } catch {
                tokenDecimals[otherToken] = 18; // 默认18位小数
            }
        }
    }

    /**
     * @dev 配置年化利率、最小质押数量和最短质押时长
     * @param _annualInterestRate 新的年化利率
     * @param _minStakeAmount 最小质押数量
     * @param _minStakeDuration 最短质押时长
     */
    function configureStakingParameters(
        uint256 _annualInterestRate,
        uint256 _minStakeAmount,
        uint256 _minStakeDuration
    ) external onlyOwner {
        annualInterestRate = _annualInterestRate;
        minStakeAmount = _minStakeAmount;
        minStakeDuration = _minStakeDuration;
    }

    /**
     * @dev 质押基础代币和其他代币
     * @param _baseAmount 基础代币数量，单位为wei
     */
    function stake(uint256 _baseAmount) external nonReentrant {
        require(
            _baseAmount >= minStakeAmount,
            "Stake amount must be greater than minimum stake amount"
        );

        StakeInfo storage stakeInfo = stakes[msg.sender];

        uint256[] memory requiredAmounts = new uint256[](otherTokens.length);

        for (uint256 i = 0; i < otherTokens.length; i++) {
            IERC20 otherToken = otherTokens[i];
            uint256 requiredAmount = (_baseAmount *
                otherTokenRatios[otherToken]) / 1e18;
            requiredAmounts[i] = requiredAmount;
            otherToken.transferFrom(msg.sender, address(this), requiredAmount);
        }

        // 转移基础代币到合约
        baseToken.transferFrom(msg.sender, address(this), _baseAmount);

        // 更新用户的质押信息
        stakeInfo.tokenStaked[baseToken] += _baseAmount;
        for (uint256 i = 0; i < otherTokens.length; i++) {
            stakeInfo.tokenStaked[otherTokens[i]] += requiredAmounts[i];
        }

        // 更新总质押量
        totalBaseStaked += _baseAmount;

        if (stakeInfo.startTime == 0) {
            stakeInfo.startTime = block.timestamp;
        }

        if (stakeInfo.lastClaimTime == 0) {
            stakeInfo.lastClaimTime = block.timestamp;
        }

        uint8 nftStatus = _determineNftStatus(stakeInfo);

        emit Staked(
            msg.sender,
            _baseAmount,
            stakeInfo.startTime,
            totalBaseStaked,
            nftStatus
        );
    }

    /**
     * @dev 解除质押，返还基础代币和其他代币
     * @param _amount 要解除质押的基础代币数量
     */
    function unstake(uint256 _amount) external nonReentrant {
        StakeInfo storage stakeInfo = stakes[msg.sender];
        require(_amount > 0, "Unstake amount must be greater than 0");
        require(
            stakeInfo.tokenStaked[baseToken] >= _amount,
            "Unstake amount exceeds staked amount"
        );
        require(
            block.timestamp >= stakeInfo.startTime + minStakeDuration,
            "Stake duration not met"
        );

        stakeInfo.tokenStaked[baseToken] -= _amount;

        for (uint256 i = 0; i < otherTokens.length; i++) {
            IERC20 token = otherTokens[i];
            uint256 unstakeOtherAmount = (_amount * otherTokenRatios[token]) /
                1e18;
            stakeInfo.tokenStaked[token] -= unstakeOtherAmount;
            token.transfer(msg.sender, unstakeOtherAmount);
        }

        // 更新总质押量
        totalBaseStaked -= _amount;

        baseToken.transfer(msg.sender, _amount);

        uint8 nftStatus = _determineNftStatus(stakeInfo);

        emit Unstaked(msg.sender, _amount, totalBaseStaked, nftStatus);
    }

    /**
     * @dev 领取奖励代币
     */
    function claimReward() external nonReentrant {
        StakeInfo storage stakeInfo = stakes[msg.sender];
        require(stakeInfo.tokenStaked[baseToken] > 0, "No staked amount");
        require(
            block.timestamp >= stakeInfo.startTime + minStakeDuration,
            "Stake duration not met"
        );

        uint256 stakingDuration = block.timestamp - stakeInfo.lastClaimTime;
        uint256 rewardAmount = _calculateReward(
            stakeInfo.tokenStaked[baseToken],
            stakingDuration
        );
        stakeInfo.lastClaimTime = block.timestamp;

        rewardToken.transfer(msg.sender, rewardAmount);

        emit RewardClaimed(msg.sender, rewardAmount, block.timestamp);

        if (
            stakeInfo.tokenStaked[baseToken] >= nftThreshold &&
            !stakeInfo.nftClaimed
        ) {
            rewardNFT.safeMint(msg.sender); // 铸造一个新的NFT给用户
            stakeInfo.nftClaimed = true;
            emit NFTRewarded(msg.sender, 0); // 假设mint返回tokenId，如果需要可以更新
        }
    }

    /**
     * @dev 计算奖励数量
     * @param _amount 质押的基础代币数量
     * @param _duration 质押时长
     */
    function _calculateReward(
        uint256 _amount,
        uint256 _duration
    ) internal view returns (uint256) {
        uint256 annualReward = (_amount * annualInterestRate) / 10000;
        uint256 reward = (annualReward * _duration) / 365 days;
        return reward;
    }

    /**
     * @dev 获取质押信息
     * @param _user 用户地址
     */
    function getStakingInfo(
        address _user
    )
        external
        view
        returns (
            uint256 startTime,
            uint256 lastClaimTime,
            uint256[] memory amounts,
            uint256 pendingReward,
            uint8 nftStatus,
            uint256 currentAnnualInterestRate
        )
    {
        StakeInfo storage stakeInfo = stakes[_user];
        uint256 tokenCount = otherTokens.length + 1; // 包含基础代币
        uint256[] memory amounts = new uint256[](tokenCount);

        amounts[0] = stakeInfo.tokenStaked[baseToken];

        for (uint256 i = 0; i < otherTokens.length; i++) {
            amounts[i + 1] = stakeInfo.tokenStaked[otherTokens[i]];
        }

        uint256 stakingDuration = block.timestamp - stakeInfo.lastClaimTime;
        uint256 rewardAmount = _calculateReward(
            stakeInfo.tokenStaked[baseToken],
            stakingDuration
        );

        uint8 nftStatus = _determineNftStatus(stakeInfo);

        return (
            stakeInfo.startTime,
            stakeInfo.lastClaimTime,
            amounts,
            rewardAmount,
            nftStatus,
            annualInterestRate
        );
    }

    /**
     * @dev 获取所有代币的地址和比例，以及年化利率，最小质押数量，最短质押时长
     */
    function getTokenInfo()
        external
        view
        returns (
            IERC20[] memory tokens,
            uint256[] memory ratios,
            uint256 currentAnnualInterestRate,
            uint256 currentNftThreshold,
            uint256 currentMinStakeDuration,
            uint256 currentMinStakeAmount,
            address rewardTokenAddress,
            address rewardNFTAddress
        )
    {
        uint256 tokenCount = otherTokens.length + 1; // 包含基础代币
        IERC20[] memory tokens = new IERC20[](tokenCount);
        uint256[] memory ratios = new uint256[](tokenCount);

        tokens[0] = baseToken;
        ratios[0] = 1e18; // 基础代币的比例为 1

        for (uint256 i = 0; i < otherTokens.length; i++) {
            tokens[i + 1] = otherTokens[i];
            ratios[i + 1] = otherTokenRatios[otherTokens[i]];
        }

        return (
            tokens,
            ratios,
            annualInterestRate,
            nftThreshold,
            minStakeDuration,
            minStakeAmount,
            address(rewardToken),
            address(rewardNFT)
        );
    }

    /**
     * @dev 提现奖励代币，仅限合约所有者使用
     * @param _amount 提现数量
     */
    function withdrawRewardTokens(uint256 _amount) external onlyOwner {
        require(_amount > 0, "Withdraw amount must be greater than 0");
        require(
            rewardToken.balanceOf(address(this)) >= _amount,
            "Withdraw amount exceeds balance"
        );

        rewardToken.transfer(msg.sender, _amount);
    }

    /**
     * @dev 确定NFT状态
     * @param stakeInfo 用户质押信息
     * @return uint8 0 表示不可领取，1 表示可领取，2 表示已领取
     */
    function _determineNftStatus(
        StakeInfo storage stakeInfo
    ) internal view returns (uint8) {
        if (stakeInfo.nftClaimed) {
            return 2;
        } else if (stakeInfo.tokenStaked[baseToken] >= nftThreshold) {
            return 1;
        } else {
            return 0;
        }
    }
}
