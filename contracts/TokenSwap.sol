// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface IERC20Detailed is IERC20 {
    function decimals() external view returns (uint8);
}

contract TokenSwap is Ownable {
    IERC20Detailed public tokenIn; // 输入代币合约（例如TokenIn）
    IERC20Detailed public tokenOut; // 输出代币合约（例如TokenOut）
    uint256 public swapRatio; // 兑换比例，以TokenIn为基准

    uint8 public tokenInDecimals; // 输入代币的精度
    uint8 public tokenOutDecimals; // 输出代币的精度

    event TokensSwapped(
        address indexed user,
        uint256 amountIn,
        uint256 amountOut
    );

    /**
     * @dev 构造函数，初始化代币合约地址和兑换比例
     * @param _tokenIn 输入代币合约地址
     * @param _tokenOut 输出代币合约地址
     * @param _swapRatio 兑换比例，以TokenIn为基准，例如1:2的比例应设置为2（表示1 TokenIn = 2 TokenOut）
     */
    constructor(
        IERC20Detailed _tokenIn,
        IERC20Detailed _tokenOut,
        uint256 _swapRatio
    ) Ownable(msg.sender) {
        require(
            address(_tokenIn) != address(_tokenOut),
            "tokenIn and tokenOut cannot be the same"
        );

        tokenIn = _tokenIn;
        tokenOut = _tokenOut;
        swapRatio = _swapRatio;

        tokenInDecimals = _getDecimals(_tokenIn);
        tokenOutDecimals = _getDecimals(_tokenOut);
    }

    /**
     * @dev 安全获取代币的精度，如果没有实现decimals方法，默认返回18
     * @param token ERC20代币合约地址
     * @return uint8 代币的精度
     */
    function _getDecimals(IERC20Detailed token) internal view returns (uint8) {
        try token.decimals() returns (uint8 decimals) {
            return decimals;
        } catch {
            return 18; // 默认精度为18
        }
    }

    /**
     * @dev 用户使用TokenIn兑换TokenOut
     * @param amountIn 兑换的TokenIn数量
     */
    function swap(uint256 amountIn) external {
        require(amountIn > 0, "Amount must be greater than zero");
        require(
            tokenIn.balanceOf(msg.sender) >= amountIn,
            "Insufficient TokenIn balance"
        );
        require(
            tokenIn.allowance(msg.sender, address(this)) >= amountIn,
            "Allowance is not enough"
        );

        uint256 amountOut;
        if (tokenOutDecimals >= tokenInDecimals) {
            uint256 decimalsDiff = uint256(tokenOutDecimals) -
                uint256(tokenInDecimals);
            amountOut = amountIn * swapRatio * (10 ** decimalsDiff);
        } else {
            uint256 decimalsDiff = uint256(tokenInDecimals) -
                uint256(tokenOutDecimals);
            amountOut = (amountIn * swapRatio) / (10 ** decimalsDiff);
        }

        require(amountOut > 0, "AmountOut must be greater than zero");
        require(
            tokenOut.balanceOf(address(this)) >= amountOut,
            "Insufficient TokenOut balance in contract"
        );

        tokenIn.transferFrom(msg.sender, address(this), amountIn);
        tokenOut.transfer(msg.sender, amountOut);

        emit TokensSwapped(msg.sender, amountIn, amountOut);
    }

    /**
     * @dev 提取合约中的TokenIn
     * @param amount 提取的TokenIn数量
     */
    function withdrawTokenIn(uint256 amount) external onlyOwner {
        tokenIn.transfer(owner(), amount);
    }

    /**
     * @dev 提取合约中的TokenOut
     * @param amount 提取的TokenOut数量
     */
    function withdrawTokenOut(uint256 amount) external onlyOwner {
        tokenOut.transfer(owner(), amount);
    }

    /**
     * @dev 设置新的兑换比例
     * @param _swapRatio 新的兑换比例
     */
    function setSwapRatio(uint256 _swapRatio) external onlyOwner {
        swapRatio = _swapRatio;
    }

    /**
     * @dev 获取tokenIn和tokenOut的合约地址及当前兑换比例
     * @return tokenInAddress tokenIn的合约地址
     * @return tokenOutAddress tokenOut的合约地址
     * @return currentSwapRatio 当前的兑换比例
     * @return tokenInDecimals 输入代币的精度
     * @return tokenOutDecimals 输出代币的精度
     */
    function getTokenInfo()
        external
        view
        returns (
            address tokenInAddress,
            address tokenOutAddress,
            uint256 currentSwapRatio,
            uint8 tokenInDecimals,
            uint8 tokenOutDecimals
        )
    {
        return (
            address(tokenIn),
            address(tokenOut),
            swapRatio,
            tokenInDecimals,
            tokenOutDecimals
        );
    }

    /**
     * @dev 获取合约中tokenIn和tokenOut的余额
     * @return tokenInBalance 合约中的tokenIn余额
     * @return tokenOutBalance 合约中的tokenOut余额
     */
    function getBalances()
        external
        view
        returns (uint256 tokenInBalance, uint256 tokenOutBalance)
    {
        tokenInBalance = tokenIn.balanceOf(address(this));
        tokenOutBalance = tokenOut.balanceOf(address(this));
    }
}
