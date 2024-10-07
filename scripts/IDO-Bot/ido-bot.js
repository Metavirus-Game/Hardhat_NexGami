const hre = require("hardhat");
const { Tokens } = require("../include/tokens")
const utils = require("../include/hre_utils");

//开始时间
const startTime = new Date("2024-06-12-21:55").getTime();
//持续时间
const duration = 60 * 10 * 1000
//金额和几率
const amounts = [
	{ amount: 100, rate: 1 },
	{ amount: 200, rate: 2 },
	{ amount: 300, rate: 1 },
	{ amount: 400, rate: 2 },
	{ amount: 500, rate: 4 },
	{ amount: 1000, rate: 8 },
	{ amount: 1500, rate: 8 },
	{ amount: 2000, rate: 8 },
	{ amount: 2100, rate: 5 },
	{ amount: 2200, rate: 5 },
	{ amount: 2500, rate: 5 },
	{ amount: 2800, rate: 5 },
	{ amount: 3000, rate: 5 },
	{ amount: 3200, rate: 5 },
	{ amount: 3500, rate: 5 },
	{ amount: 3800, rate: 3 },
	{ amount: 4000, rate: 3 },
	{ amount: 4100, rate: 3 },
	{ amount: 4200, rate: 3 },
	{ amount: 4500, rate: 3 },
	{ amount: 4700, rate: 3 },
	{ amount: 5000, rate: 3 },
	{ amount: 5500, rate: 2 },
	{ amount: 6000, rate: 2 },
	{ amount: 6100, rate: 2 },
	{ amount: 6200, rate: 1 },
	{ amount: 6300, rate: 1 },
	{ amount: 6500, rate: 1 },
	{ amount: 6800, rate: 1 },
	{ amount: 7000, rate: 1 },
	{ amount: 7500, rate: 1 },
	{ amount: 8000, rate: 1 },
	{ amount: 8200, rate: 1 },
	{ amount: 8500, rate: 1 },
	{ amount: 8600, rate: 1 },
	{ amount: 8800, rate: 1 },
	{ amount: 10000, rate: 1 },
]


const tremble = {
	/**
	 * 抖动几率 0-100
	 */
	prob: 5,

	/**
	 * 抖动范围
	 */
	range: { min: -50, max: 50 }
}

/**
 * 希望的充值总数
 * 充值1200万USD
 */
const expectingAmount = 12000000;

class User {
	chargeAmount;
	address;
}

/**
 * @type {User[]}
 */
const users = [];

const main = async () => {
	const signers = await hre.ethers.getSigners();
	console.log(signers[2])
	randomUsers();
}

const randomUsers = () => {
	let total = 0;
	const totalWeight = amounts.reduce((s, u) => s + u.rate, 0)

	while (total < expectingAmount) {
		const w = Math.random() * totalWeight;
		let idx = 0;
		let t = 0;
		while (t < w) {
			t += amounts[idx].rate;
			if (idx == amounts.length - 1) break;
			idx++;
		}
		const u = new User();
		u.chargeAmount = randomAmount(amounts[idx].amount);
		total += u.chargeAmount
		users.push(u);
	}


	console.log("一共", users.length, "人");

	const o = [];

	for (const idx in users) {
		const amt = users[idx].chargeAmount;

		let key = amt;


		if (o[key] == null) {
			o[key] = 1;
		} else {
			o[key]++
		}
	}

	for (const idx in o) {
		if (o[idx] != null) {
			console.log(`额度\t${idx}\t${o[idx]}人`)
		}
	}
}

const randomAmount = (amount) => {
	const r = Math.random() * 100;
	let ret = amount;
	if (r < tremble.prob) {
		const t = (tremble.range.max - tremble.range.min) * Math.random() + tremble.range.min;
		ret += parseInt(t);
		if (ret > 10000) ret = 10000;//最多就冲10000
	}
	return ret;
}

main().catch(e => {
	console.error(e);
	process.exit(1);
})