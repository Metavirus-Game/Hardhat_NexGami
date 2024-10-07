// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MultiTokenStakingWithLock is Ownable {
    // 质押信息结构体
    struct Stake {
        uint256 amount; // 基础代币质押数量
        uint256 rewardDebt; // 奖励债务
        uint256 unlockTime; // 解锁时间戳
    }

    // 代币信息结构体
    struct TokenInfo {
        IERC20 token; // 代币合约
        uint256 ratio; // 相对于基础代币的比例
        uint256 totalStaked; // 总质押数量
    }

    IERC20 public baseToken; // 基础质押代币
    uint256 public rewardRate = 1e18; // 每秒奖励代币数量
    uint256 public rewardPerTokenStored; // 每个基础代币的累积奖励
    uint256 public lastUpdateTime; // 上次更新奖励的时间戳
    uint256 public totalBaseStaked; // 总基础代币质押数量
    uint256 public lockPeriod = 30 days; // 锁定期（秒为单位）

    mapping(address => Stake) public stakes; // 存储每个用户的质押信息
    TokenInfo[] public tokenInfos; // 存储所有代币的信息
    mapping(IERC20 => uint256) public tokenIndex; // 存储每个代币在tokenInfos中的索引
    mapping(address => uint256) public rewards; // 存储每个用户的奖励数量

    event Staked(address indexed user, uint256 amount); // 质押事件
    event Withdrawn(address indexed user, uint256 amount); // 提现事件
    event RewardPaid(address indexed user, uint256 reward); // 奖励支付事件

    modifier updateReward(address account) {
        rewardPerTokenStored = rewardPerToken();
        lastUpdateTime = block.timestamp;

        if (account != address(0)) {
            rewards[account] = earned(account);
            stakes[account].rewardDebt = rewardPerTokenStored;
        }
        _;
    }

    // 构造函数，初始化基础代币合约和最后更新时间
    constructor(IERC20 _baseToken) Ownable(msg.sender) {
        baseToken = _baseToken;
        lastUpdateTime = block.timestamp;
    }

    /**
     * @dev 添加新的支持质押的代币及其比例
     * @param token 代币合约地址
     * @param ratio 代币相对于基础代币的比例
     */
    function addToken(IERC20 token, uint256 ratio) external onlyOwner {
        require(
            tokenIndex[token] == 0 &&
                (tokenInfos.length == 0 || tokenInfos[0].token != token),
            "Token already added"
        );

        tokenInfos.push(
            TokenInfo({token: token, ratio: ratio, totalStaked: 0})
        );
        tokenIndex[token] = tokenInfos.length - 1;
    }

    /**
     * @dev 质押基础代币和其他代币
     * @param amount 质押的基础代币数量
     */
    function stake(uint256 amount) external updateReward(msg.sender) {
        require(amount > 0, "Cannot stake 0");

        // 质押基础代币
        baseToken.transferFrom(msg.sender, address(this), amount);
        stakes[msg.sender].amount += amount;
        totalBaseStaked += amount;

        // 按比例质押其他代币
        for (uint256 i = 0; i < tokenInfos.length; i++) {
            TokenInfo storage info = tokenInfos[i];
            uint256 tokenAmount = (amount * info.ratio) / 1e18;
            info.token.transferFrom(msg.sender, address(this), tokenAmount);
            info.totalStaked += tokenAmount;
        }

        stakes[msg.sender].unlockTime = block.timestamp + lockPeriod;

        emit Staked(msg.sender, amount);
    }

    /**
     * @dev 提现质押的代币
     * @param amount 提现的基础代币数量
     */
    function withdraw(uint256 amount) public updateReward(msg.sender) {
        require(amount > 0, "Cannot withdraw 0");
        require(
            stakes[msg.sender].amount >= amount,
            "Withdraw amount exceeds balance"
        );
        require(
            block.timestamp >= stakes[msg.sender].unlockTime,
            "Tokens are still locked"
        );

        stakes[msg.sender].amount -= amount;
        totalBaseStaked -= amount;
        baseToken.transfer(msg.sender, amount);

        // 按比例提现其他代币
        for (uint256 i = 0; i < tokenInfos.length; i++) {
            TokenInfo storage info = tokenInfos[i];
            uint256 tokenAmount = (amount * info.ratio) / 1e18;
            info.token.transfer(msg.sender, tokenAmount);
            info.totalStaked -= tokenAmount;
        }

        emit Withdrawn(msg.sender, amount);
    }

    /**
     * @dev 获取用户的奖励
     */
    function getReward() public updateReward(msg.sender) {
        uint256 reward = rewards[msg.sender];
        if (reward > 0) {
            rewards[msg.sender] = 0;
            baseToken.transfer(msg.sender, reward);
            emit RewardPaid(msg.sender, reward);
        }
    }

    /**
     * @dev 用户退出质押
     */
    function exit() external {
        withdraw(stakes[msg.sender].amount);
        getReward();
    }

    /**
     * @dev 计算每个基础代币应得的奖励
     * @return uint256 每个基础代币应得的奖励
     */
    function rewardPerToken() public view returns (uint256) {
        if (totalBaseStaked == 0) {
            return rewardPerTokenStored;
        }
        return
            rewardPerTokenStored +
            (((block.timestamp - lastUpdateTime) * rewardRate * 1e18) /
                totalBaseStaked);
    }

    /**
     * @dev 计算特定用户的应得奖励
     * @param account 用户地址
     * @return uint256 用户应得的奖励
     */
    function earned(address account) public view returns (uint256) {
        return
            ((stakes[account].amount *
                (rewardPerToken() - stakes[account].rewardDebt)) / 1e18) +
            rewards[account];
    }

    /**
     * @dev 设置奖励率
     * @param _rewardRate 新的奖励率
     */
    function setRewardRate(uint256 _rewardRate) external onlyOwner {
        rewardRate = _rewardRate;
    }

    /**
     * @dev 设置锁定期时间
     * @param _lockPeriod 新的锁定期时间（秒为单位）
     */
    function setLockPeriod(uint256 _lockPeriod) external onlyOwner {
        lockPeriod = _lockPeriod;
    }

    /**
     * @dev 为合约添加奖励代币
     * @param amount 奖励代币数量
     */
    function addRewardTokens(uint256 amount) external onlyOwner {
        baseToken.transferFrom(msg.sender, address(this), amount);
    }
}
