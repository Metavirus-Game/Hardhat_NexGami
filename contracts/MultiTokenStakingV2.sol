// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./includes/IERC20Extended.sol";
import "./includes/IERC721.sol";

/**
 * @title MultiTokenStakingV2
 * @dev 允许用户质押基础代币和其他代币，在质押成功时发放NFT，claim时返还本金并按比例返还利息。
 */
contract MultiTokenStakingV2 is Ownable, ReentrancyGuard {
    struct StakeInfo {
        uint256 startTime; // 质押开始时间
        uint256 lastClaimTime; // 上次领取奖励的时间
        mapping(IERC20 => uint256) tokenStaked; // 用户质押的代币数量，包括基础代币和其他代币
    }

    IERC20Extended public baseToken; // 基础代币
    IERC721 public rewardNFT; // 奖励NFT
    uint256 public returnRate; // 返还比例，单位为百分之一，例如返还比例2%则为200
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
        uint256 totalBaseStaked
    ); // 质押事件
    event Claimed(
        address indexed user,
        uint256 principal,
        uint256 interest,
        uint256 timestamp
    ); // 领取事件
    event NFTRewarded(address indexed user, uint256 tokenId); // NFT奖励事件

    /**
     * @dev 构造函数，初始化合约
     * @param _baseToken 基础代币地址
     * @param _returnRate 返还比例
     * @param _minStakeDuration 获得奖励的最短质押时长
     * @param _minStakeAmount 最小质押数量
     * @param _rewardNFT 奖励NFT地址
     */
    constructor(
        address _baseToken,
        uint256 _returnRate,
        uint256 _minStakeDuration,
        uint256 _minStakeAmount,
        address _rewardNFT
    ) Ownable(msg.sender) {
        baseToken = IERC20Extended(_baseToken);
        returnRate = _returnRate;
        rewardNFT = IERC721(_rewardNFT);
        minStakeDuration = _minStakeDuration;
        minStakeAmount = _minStakeAmount;

        try baseToken.decimals() returns (uint8 decimals) {
            tokenDecimals[baseToken] = decimals;
        } catch {
            tokenDecimals[baseToken] = 18; // 默认18位小数
        }
    }

    /**
     * @dev 配置质押代币和比例
     * @param _otherTokens 其他代币数组
     * @param _ratios 其他代币与基础代币的比例数组，单位为wei
     */
    function configureStaking(
        IERC20[] calldata _otherTokens,
        uint256[] calldata _ratios
    ) external onlyOwner {
        require(
            _otherTokens.length == _ratios.length,
            "Token and ratio length mismatch"
        );

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
     * @dev 配置返还比例、最小质押数量和最短质押时长
     * @param _returnRate 新的返还比例
     * @param _minStakeAmount 最小质押数量
     * @param _minStakeDuration 最短质押时长
     */
    function configureStakingParameters(
        uint256 _returnRate,
        uint256 _minStakeAmount,
        uint256 _minStakeDuration
    ) external onlyOwner {
        returnRate = _returnRate;
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

        // 发放NFT奖励
        rewardNFT.safeMint(msg.sender);
        emit NFTRewarded(msg.sender, 0); // 假设mint返回tokenId，如果需要可以更新

        emit Staked(
            msg.sender,
            _baseAmount,
            stakeInfo.startTime,
            totalBaseStaked
        );
    }

    /**
     * @dev 领取本金和利息
     */
    function claim() external nonReentrant {
        StakeInfo storage stakeInfo = stakes[msg.sender];
        require(stakeInfo.tokenStaked[baseToken] > 0, "No staked amount");
        require(
            block.timestamp >= stakeInfo.startTime + minStakeDuration,
            "Stake duration not met"
        );

        uint256 principal = stakeInfo.tokenStaked[baseToken];
        uint256 interest = (principal * returnRate) / 10000; // 按比例计算利息

        stakeInfo.tokenStaked[baseToken] = 0;

        // 返还本金和利息
        baseToken.transfer(msg.sender, principal + interest);

        // 返还其他代币
        for (uint256 i = 0; i < otherTokens.length; i++) {
            IERC20 otherToken = otherTokens[i];
            uint256 otherPrincipal = stakeInfo.tokenStaked[otherToken];
            uint256 otherInterest = (otherPrincipal * returnRate) / 10000;
            stakeInfo.tokenStaked[otherToken] = 0;
            otherToken.transfer(msg.sender, otherPrincipal + otherInterest);
        }

        emit Claimed(msg.sender, principal, interest, block.timestamp);
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
            uint256 pendingInterest,
            uint256 currentReturnRate
        )
    {
        StakeInfo storage stakeInfo = stakes[_user];
        uint256 tokenCount = otherTokens.length + 1; // 包含基础代币
        uint256[] memory amounts = new uint256[](tokenCount);

        amounts[0] = stakeInfo.tokenStaked[baseToken];

        for (uint256 i = 0; i < otherTokens.length; i++) {
            amounts[i + 1] = stakeInfo.tokenStaked[otherTokens[i]];
        }

        uint256 pendingInterest = (stakeInfo.tokenStaked[baseToken] *
            returnRate) / 10000;

        return (
            stakeInfo.startTime,
            stakeInfo.lastClaimTime,
            amounts,
            pendingInterest,
            returnRate
        );
    }

    /**
     * @dev 获取所有代币的地址和比例，以及返还比例，最小质押数量，最短质押时长
     */
    function getTokenInfo()
        external
        view
        returns (
            IERC20[] memory tokens,
            uint256[] memory ratios,
            uint256 currentReturnRate,
            uint256 currentMinStakeDuration,
            uint256 currentMinStakeAmount,
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
            returnRate,
            minStakeDuration,
            minStakeAmount,
            address(rewardNFT)
        );
    }

    /**
     * @dev 提现代币，仅限合约所有者使用
     * @param _baseAmount 提现基础代币的数量
     */
    function withdrawTokens(uint256 _baseAmount) external onlyOwner {
        require(_baseAmount > 0, "Withdraw amount must be greater than 0");
        uint256 baseBalance = baseToken.balanceOf(address(this));
        if (_baseAmount > baseBalance) {
            _baseAmount = baseBalance;
        }

        // 先检查所有代币余额是否足够
        for (uint256 i = 0; i < otherTokens.length; i++) {
            IERC20 otherToken = otherTokens[i];
            uint256 otherAmount = (_baseAmount * otherTokenRatios[otherToken]) /
                1e18;
            uint256 otherBalance = otherToken.balanceOf(address(this));
            if (otherAmount > otherBalance) {
                otherAmount = otherBalance;
            }
            require(
                otherToken.balanceOf(address(this)) >= otherAmount,
                "Withdraw amount exceeds balance for other tokens"
            );
        }

        // 提取基础代币
        baseToken.transfer(msg.sender, _baseAmount);

        // 提取其他代币
        for (uint256 i = 0; i < otherTokens.length; i++) {
            IERC20 otherToken = otherTokens[i];
            uint256 otherAmount = (_baseAmount * otherTokenRatios[otherToken]) /
                1e18;
            otherToken.transfer(msg.sender, otherAmount);
        }
    }
}
