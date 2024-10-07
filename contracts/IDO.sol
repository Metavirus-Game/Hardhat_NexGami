// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract IDO is Ownable {
    IERC20 public saleToken; // 代币合约
    uint256 public tokenPrice; // 代币价格（以ETH为单位）
    uint256 public totalTokens; // 代币总量
    uint256 public tokensSold; // 已售出的代币数量
    uint256 public startTime; // 销售开始时间
    uint256 public endTime; // 销售结束时间
    bool public emergencyStop; // 紧急停止开关
    IERC20 public paymentToken; // 支付代币合约（如果使用ERC20代币）

    mapping(address => uint256) public purchases; // 记录每个参与者购买的代币数量

    event TokensPurchased(
        address indexed purchaser,
        uint256 amount,
        uint256 value,
        address paymentToken
    );
    event TokensWithdrawn(address indexed beneficiary, uint256 amount);
    event TokensRefunded(address indexed user, uint256 amount);

    modifier onlyDuringSale() {
        require(
            block.timestamp >= startTime && block.timestamp <= endTime,
            "Sale is not active"
        );
        require(!emergencyStop, "Sale is stopped");
        _;
    }

    constructor(
        IERC20 _saleToken,
        uint256 _tokenPrice,
        uint256 _totalTokens,
        uint256 _startTime,
        uint256 _endTime,
        IERC20 _paymentToken
    ) Ownable(msg.sender) {
        require(_startTime < _endTime, "Start time must be before end time");
        require(_totalTokens > 0, "Total tokens must be greater than zero");

        saleToken = _saleToken;
        tokenPrice = _tokenPrice;
        totalTokens = _totalTokens;
        startTime = _startTime;
        endTime = _endTime;
        emergencyStop = false;
        paymentToken = _paymentToken; // 如果不使用ERC20代币支付，可以设置为address(0)
    }

    /**
     * @dev 参与者购买代币
     */
    function buyTokens(
        uint256 amount,
        bool useERC20
    ) external payable onlyDuringSale {
        uint256 tokens;
        uint256 value;

        if (useERC20) {
            require(
                address(paymentToken) != address(0),
                "ERC20 payment token not set"
            );
            value = amount * tokenPrice;
            paymentToken.transferFrom(msg.sender, address(this), value);
            tokens = amount;
        } else {
            value = msg.value;
            require(value > 0, "ETH amount must be greater than zero");
            tokens = value / tokenPrice;
            require(tokens * tokenPrice == value, "Incorrect ETH amount sent");
        }

        require(
            tokensSold + tokens <= totalTokens,
            "Not enough tokens left for sale"
        );

        purchases[msg.sender] += tokens;
        tokensSold += tokens;

        emit TokensPurchased(
            msg.sender,
            tokens,
            value,
            useERC20 ? address(paymentToken) : address(0)
        );
    }

    /**
     * @dev 项目团队提取ETH或ERC20支付的资金
     */
    function withdrawFunds() external onlyOwner {
        require(
            block.timestamp > endTime || emergencyStop,
            "Sale has not ended"
        );
        if (address(paymentToken) != address(0)) {
            uint256 balance = paymentToken.balanceOf(address(this));
            paymentToken.transfer(owner(), balance);
        }
        uint256 ethBalance = address(this).balance;
        if (ethBalance > 0) {
            payable(owner()).transfer(ethBalance);
        }
    }

    /**
     * @dev 紧急停止IDO
     */
    function stopIDO() external onlyOwner {
        emergencyStop = true;
    }

    /**
     * @dev 恢复IDO
     */
    function resumeIDO() external onlyOwner {
        emergencyStop = false;
    }

    /**
     * @dev 提取未售出的代币
     */
    function withdrawUnsoldTokens() external onlyOwner {
        require(
            block.timestamp > endTime || emergencyStop,
            "Sale has not ended"
        );

        uint256 unsoldTokens = totalTokens - tokensSold;
        require(unsoldTokens > 0, "No unsold tokens to withdraw");

        saleToken.transfer(owner(), unsoldTokens);

        emit TokensWithdrawn(owner(), unsoldTokens);
    }

    /**
     * @dev 用户在IDO结束后退回未售出的代币
     */
    function refundTokens() external {
        require(
            block.timestamp > endTime || emergencyStop,
            "Sale has not ended"
        );

        uint256 purchasedAmount = purchases[msg.sender];
        require(purchasedAmount > 0, "No purchased tokens to refund");

        uint256 refundAmount = purchasedAmount * tokenPrice;
        purchases[msg.sender] = 0;

        if (address(paymentToken) != address(0)) {
            paymentToken.transfer(msg.sender, refundAmount);
        } else {
            payable(msg.sender).transfer(refundAmount);
        }

        emit TokensRefunded(msg.sender, refundAmount);
    }

    /**
     * @dev 查看代币销售进度
     * @return total 总代币数量
     * @return sold 已售出代币数量
     * @return start 销售开始时间
     * @return end 销售结束时间
     */
    function getSaleProgress()
        external
        view
        returns (uint256 total, uint256 sold, uint256 start, uint256 end)
    {
        total = totalTokens;
        sold = tokensSold;
        start = startTime;
        end = endTime;
    }
}
