const hre = require("hardhat");
const { Tokens, TokenDef } = require("./include/tokens.js");
const { toSecond, waitConfirm, waitTransaction, tokenBalanceOf } = require("./include/hre_utils.js");
const colors = require("colors-console");

// 释放间隔定义
const ReleaseInterval = {
	Minute: 0,
	Hour: 1,
	Day: 2,
	Week: 3,
	Month: 4,
	Quarter: 5,
	HalfYear: 6,
	Year: 7
}

const StrReleaseInterval = [
	"每分钟", "每小时", "每天", "每周", "每月", "每季度", "每半年", "每年"
]

class VestingOption {
	/**
	 * 受益人地址
	 * @type {string}
	 */
	beneficiary;
	/**
	 * 发放的 token 
	 * @type {TokenDef}
	 */
	token;

	/**
	 * 发放数量，单位ether
	 * @type {string}
	 */
	amount;

	/**
	 * 开始时间
	 * @type {Date}
	 */
	start;

	/**
	 * 归属时长，单位秒
	 * @type {number}
	 */
	duration;

	/**
	 * 受益人角色，0=investor 1=employee
	 * @type {0|1}
	 */
	role;

	/**
	 * 释放周期
	 * @type {number}
	 * @see {ReleaseInterval}
	 */
	releaseInterval;

	/**
	 * 悬崖期，单位秒
	 * @type {number}
	 */
	cliff;

	/**
	 * 悬崖期到期后释放比例，0~10000，2000表示20%
	 */
	cliffReleasePercentage;
}

// const contractAddress = "0x2815Ec3083f47f073DE9EB5670CfCbF1BD1709dc" //MVT
const contractAddress = {
	polygon_amoy: "0xdF8167D00f5ca0AFF1FcEb2e6427D2C79e8CEe60", //NEXG
	polygon_mainnet: "0xA85B4D49CB2963f1CEB2A68290B680043aF31089"
}

const getVestingContract = () => {
	return contractAddress[hre.network.name]
}

async function main() {
	// await deploy(Tokens.get("MVT"));
	// await deploy(Tokens.get("NEXG"));
	//await deploy(Tokens.get("NEXG"))
	await createVesting({
		beneficiary: "0x7758c5403094CAA164840aEf70561ceee0CcAb8A",
		token: Tokens.get("NEXG"),
		amount: "111111.5",
		start: new Date("2024-06-30 10:00:00"),
		duration: toSecond(3, "Month"),
		role: 0,
		releaseInterval: ReleaseInterval.Day,
		cliff: 0,
		cliffReleasePercentage: 0,
	});
	//await createTestVesting();
	//await getVestingSchedule();
}

async function deploy(vestingToken) {
	if (vestingToken == null) {
		throw "Token Config Error " + hre.network.name
	}
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy TokenVesting for [${vestingToken.name}]`)
	const myth = await hre.ethers.deployContract("TokenVesting", [vestingToken.getAddress()], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`TokenVesting deployed to ${myth.target}`
	);
}


/**
 * @param {VestingOption} option 设置
 */
async function createVesting(option) {
	const [owner] = await ethers.getSigners();
	const con = await hre.ethers.getContractAt("TokenVesting", getVestingContract(), owner);
	const start = Math.floor(option.start / 1000);
	const cliff = option.cliff;
	const duration = option.duration;
	const totalAmount = ethers.parseUnits(option.amount, option.token.decimals);//ethers.parseUnits("2500000", 18); // 归属总量为1000代币
	const revocable = option.role == 1; // 归属计划可撤销
	const role = option.role;
	const releaseInterval = option.releaseInterval;
	const cliffReleasePercentage = option.cliffReleasePercentage;

	console.group(colors(['green', 'bright'], `Creating Token Vesting on [${hre.network.name}] Contract[${getVestingContract()}]`));
	console.log(`Beneficiary:\t${role == 0 ? 'Investor' : 'Employee'} -- ${option.beneficiary}`);
	console.log(`Token:\t${option.token.name} -- ${option.token.getAddress()}`);
	console.log(`Allocation:\t${option.amount} ${option.token.name} (${totalAmount})`);
	console.log(`Start Time:\t${option.start.toLocaleString()}(${start})`);
	console.log(`Duration: \t${duration}秒 -- to date ${new Date(option.start.getTime() + duration * 1000).toLocaleString()}`);
	console.log(`Cliff:\t${cliff}秒 -- to date  ${new Date(option.start.getTime() + cliff * 1000).toLocaleString()} Release ${(cliffReleasePercentage / 100.0).toFixed(2)}%`);
	console.log(`Revocable:\t`, colors(revocable ? 'green' : 'red', revocable));
	console.log(`Cadence:\t${StrReleaseInterval[releaseInterval]}`);
	console.groupEnd();
	const c = await waitConfirm();
	if (c) {
		const balance = await tokenBalanceOf(option.token.getAddress(), owner.address);
		if (balance < totalAmount) {
			console.error(`Insufficient ${option.token.name} balance of Owner[${owner.address}]`)
			console.error(`Need ${option.amount}, Current Balance: ${ethers.formatEther(balance)}`)
			process.exit(1);
		}
		const tx = await con.createVestingSchedule(
			option.beneficiary,
			start,
			cliff,
			duration,
			totalAmount,
			revocable,
			role,
			releaseInterval,
			cliffReleasePercentage
		);
		await waitTransaction(tx.hash);

		//transfer token to contract
		console.log(`transfer ${option.amount} ${option.token.name} to ${getVestingContract()}`)
		const conToken = await hre.ethers.getContractAt(option.token.contractName, option.token.getAddress(), owner);
		const tx1 = await conToken.transfer(getVestingContract(), totalAmount);
		await waitTransaction(tx1.hash)
	}
	process.exit(0);



}

async function createTestVesting() {
	const [owner] = await ethers.getSigners();
	const con = await hre.ethers.getContractAt("TokenVesting", getVestingContract(), owner);
	// const beneficiary = "0x6A0E43e2dDe3D180FFBF707aaC063b36300932D2"; //gary
	const beneficiary = "0xaDF8482D274Aa1Bbe1BC42AD4908aA1eA7B89cc4"; //gary 1
	// const beneficiary = "0x84ff9a383d39ccfd3bef81929ca5eaccef205e70";
	// const beneficiary = "0xf1029c9f37134036ffd236ebc8051d3d94178175";
	const start = Math.floor(Date.now() / 1000); // 以当前时间作为开始时间
	const cliff = 60 * 60 * 1; // 悬崖期
	const duration = 60 * 60 * 24 * 180; // 归属期为半年
	const totalAmount = ethers.parseUnits("2500000", 18); // 归属总量为1000代币
	const revocable = false; // 归属计划可撤销
	const role = 1; // 角色为1=Employee 0=investor
	const releaseInterval = 1; // 释放间隔为0=minute 1=hour 2=...
	const cliffReleasePercentage = 1000;//cliff到期后释放比例，0~10000，2000表示20%

	const tx = await con.createVestingSchedule(
		beneficiary,
		start,
		cliff,
		duration,
		totalAmount,
		revocable,
		role,
		releaseInterval,
		cliffReleasePercentage
	);

	console.log("tx:", tx)
}

async function getVestingSchedule() {
	const [owner] = await ethers.getSigners();
	const con = await hre.ethers.getContractAt("TokenVesting", getVestingContract(), owner);
	const beneficiary = "0x84ff9a383d39ccfd3bef81929ca5eaccef205e70";
	const ret = await con.getVestingSchedule(beneficiary)
	console.log(ret);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})