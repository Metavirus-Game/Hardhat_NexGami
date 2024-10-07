const hre = require("hardhat");
const ERC20 = require("../../ERC20.json");

async function tokenBalanceOf(token_address, walletAddress) {
	const provider = hre.ethers.provider;
	const token = new hre.ethers.Contract(token_address, ERC20.abi, provider);
	const balance = await token.balanceOf(walletAddress);
	return balance;
}


async function approveToken(token_address, owner, spender, amount) {
	const provider = ethers.provider;
	const usdcCon = new hre.ethers.Contract(token_address, ERC20.abi, owner);
	const tx = await usdcCon.approve(spender, amount);
	return tx.hash;
}

async function allowanceToken(token_address, owner, spender) {
	const provider = hre.ethers.provider;
	const usdcCon = new hre.ethers.Contract(token_address, ERC20.abi, owner);
	const allow = await usdcCon.allowance(owner.address, spender);
	return allow
}

async function waitTransaction(txHash) {
	let t = 0;
	process.stdout.write(`\n\rwaiting for tx [${txHash}]`)
	while (true) {
		const r = await hre.ethers.provider.getTransactionReceipt(txHash);
		process.stdout.write(".");
		if (r == null) {
			await sleep(1000);
			t += 1000;
			if (t > 5 * 60 * 1000) {
				//超时5分钟
				return false;
			}
		} else {
			process.stdout.write("Done\n\r");
			return true;
		}
	}
}

const sleep = async (ms) => {
	return new Promise((r) => {
		setTimeout(() => {
			r();
		}, ms);
	})
}

async function verify(address, contract, args) {
	console.log("verifying contract...", address);
	const result = await hre.run("verify:verify", {
		address: address,
		constructorArguments: args,
		contract: contract
	})
	console.log(result);
}


const DurationUnit = {
	Minute: 0,
	Hour: 1,
	Day: 2,
	Week: 3,
	Month: 4,
	Year: 5
}

/**
 * 将时长转化为秒数
 * @param {string} duration 
 * @param {"Minute"|"Hour"|"Day"|"Week"|"Month"|"Year"} unit 
 * @see {DurationUnit}
 */
function toSecond(duration, unit) {
	let u = 1;
	switch (unit) {
		case "Minute":
			u = 60;
			break;
		case "Hour":
			u = 60 * 60;
			break;
		case "Day":
			u = 60 * 60 * 24;
			break;
		case "Week":
			u = 60 * 60 * 24 * 7
			break;
		case "Month":
			//1个月统一按30天算
			u = 60 * 60 * 24 * 30
			break;
		case "Year":
			u = 60 * 60 * 24 * 365
			break;
	}
	return duration * u;
}

/**
 * 
 * @param {string} message 
 */
const waitConfirm = async (message) => {
	if (message != null && message.length > 0) {
		process.stdout.write(message + "  ")
	}
	process.stdout.write(`Continue[y/n]?`)

	let resp = null;

	process.stdin.once('data', (input) => {
		const a = input.toString().trim();
		if (['Y', 'y', 'YES', 'yes'].indexOf(a) > -1) {
			resp = true;
		} else {
			resp = false
		}
	})

	while (resp == null) {
		await sleep(100);
	}

	console.log("");

	return resp;

}

module.exports = { tokenBalanceOf, approveToken, allowanceToken, waitTransaction, sleep, verify, toSecond, waitConfirm, DurationUnit }