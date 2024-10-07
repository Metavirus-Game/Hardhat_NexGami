const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("TokenSwap", () => {
    let tokenIn = "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"; //usdt
    let tokenInDecimals = 6;
    let tokenOut = "0xEb75F3952273B97Aa98fe85a3dD447BC34D4B4De"; //nexu
    let tokenOutDecimals = 18;
    let tokenSwapContract = "0x01EaAA0e78156e161ddE27946C6923ddA25E25C5"; //tokenSwapPermit
    let signer;

    beforeEach(async () => {
        [_, signer] = await ethers.getSigners();
        tokenSwapContract = await hre.ethers.getContractAt("TokenSwapPermit", tokenSwapContract, signer);
        tokenIn = await ethers.getContractAt("contracts/IERC20Permit.sol:IERC20Permit", tokenIn, signer)
        console.log(signer)
    });

    it("Should swap tokens correctly", async () => {
        const amountIn = ethers.parseUnits("10", tokenInDecimals);
        const deadline = Math.floor(Date.now() / 1000) + 1800; // half an hour from now
        const nonce = await tokenIn.nonces(signer.address);

        console.log(nonce)
    })

})