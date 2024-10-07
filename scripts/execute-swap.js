const hre = require("hardhat");
const ERC20 = require("../ERC20.json");

const { swapAddrs } = require("./include/swap_addresses")
const { tokenBalanceOf, approveToken, allowanceToken, waitTransaction } = require("./include/hre_utils")

let owner = null;

const main = async () => {
	// const [_, o] = await ethers.getSigners();
	// owner = o;

	// const balanceIn = await tokenBalanceOf(swapAddrs.tokenIn, owner.address)
	// console.log(`operator [${owner.address}] has ${hre.ethers.formatUnits(balanceIn, swapAddrs.tokenInDecimals)} tokenIn`)

	// const balanceOut = await tokenBalanceOf(swapAddrs.tokenOut, swapAddrs.contractAddr)
	// console.log(`SwapContract [${swapAddrs.contractAddr}] has ${hre.ethers.formatUnits(balanceOut, swapAddrs.tokenOutDecimals)} tokenOut`)

	// await swap(1000);

	const [o] = await ethers.getSigners();
	owner = o;

	await withdraw();
}

/**
 * 
 * @param {number} amount 交换数量，最大单位(ether)
 */
const swap = async (amount) => {
	const _in = amount;
	const _out = _in * swapAddrs.ratio
	const inAmt = hre.ethers.parseUnits(_in.toString(), swapAddrs.tokenInDecimals)
	const outAmt = hre.ethers.parseUnits(_out.toString(), swapAddrs.tokenOutDecimals)

	console.log(`Swap ${inAmt}(${_in}) TokenIn to ${outAmt}(${_out}) TokenOut`)
	const txApprove = await approveToken(swapAddrs.tokenIn, owner, swapAddrs.contractAddr, inAmt)
	console.log('approving...', txApprove)
	const r = await waitTransaction(txApprove);
	if (!r) {
		console.error("approve failed");
		process.exit(1);
	}

	const swapContract = await hre.ethers.getContractAt("TokenSwap", swapAddrs.contractAddr, owner);
	const tx = await swapContract.swap(inAmt);

	const r1 = await waitTransaction(tx.hash);
	if (!r1) {
		console.error("swap failed");
		process.exit(1);
	}

	const b = await swapContract.getBalances();
	console.log(b);
}

const withdraw = async () => {
	const swapContract = await hre.ethers.getContractAt("TokenSwap", swapAddrs.contractAddr, owner);
	const r = await swapContract.getBalances();
	console.log(r);
	const tx1 = await swapContract.withdrawTokenIn(r[0])
	console.log(tx1.hash);
	const tx2 = await swapContract.withdrawTokenOut(r[1])
	console.log(tx2.hash);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})
