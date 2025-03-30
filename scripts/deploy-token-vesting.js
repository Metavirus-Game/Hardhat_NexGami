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


const NEXGVestings = [
	// { name: "HTX Ventures", address: "0x8692c98c6dCea83D4447fA2F795cAc376e8B4BDa", amount: "6000000.3", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "OIG", address: "0x1074045B6ee28F705c75fe8FFf21A15AF76d4d13", amount: "1125000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Polygon", address: "0x27ee3d1f71f41dead9661ea092f236b4bd14f380", amount: "7500000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Henry", address: "0x72146fA646f758eAF6c0C86FA64578B5B36520cc", amount: "10000000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Metalpha", address: "0x198143320acec70699cfd281dfe62625a41be081", amount: "250000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Louis", address: "0xd3101FFe033E187fAD7C81b4CdF4FDA380152698", amount: "10000000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "木头", address: "0xCD12C0118404328452316be95526aAD0989A650E", amount: "10000000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Edward", address: "0x09bEAF14FA1279be81bA30FF83c76FDba02b709B", amount: "10000000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Gate Labs", address: "0x0d35428F4C48aAb788d512dC02bB2b6Aa008dFe8", amount: "2500000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Shin", address: "0xe41b22f8916B09b40E21471aB24C0721c5aC79F9", amount: "1000000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "Kekkai", address: "0xb668C33D2498B62966dE8f7Bf8ABb6EE5C977293", amount: "50000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Capital", address: "0x3E3cF9391A7be6BC549059F1313374bb30bf6A4b", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Captial", address: "0xF140284F6f8EE7a2E4601F60eCC9d456E9724501", amount: "84037.5", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Captial", address: "0x51b117C28945437AaAF8AAF52A056e74f4974D0A", amount: "30375", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Captial", address: "0xf76d56F795da5c34050d61f134de0e87B8c41763", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Captial", address: "0xaD82aB762EC17460e72bb1EcCB98dD54101012a7", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	// { name: "WM Captial", address: "0x16Bf2e750dfD4656fd6cB76776E595111d8EeCE4", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x331Ba849507c87B345844d39a4c3f67E96b83b0a", amount: "81000", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x39c71979C935C492FAd8d09BfBF0DBE545761a51", amount: "25312.5", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x49D593cc2c265cDdc98344280482b035087B71Cc", amount: "40500", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x9B152454Da451C4AE1dcBa976c14eaCdc8Cc30F1", amount: "151875", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0xf0D4B6297c1Bd35249d99EC17efD1C8B5DCD1F19", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0xDAcD4Fcae0a9acc4197e09e551b9557Fe2d5D91d", amount: "60750", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x7248C0CB46bC8676aDc2139438A4854B674C4075", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x9aBB5b20520384E2129534ea2004AcAaEc8C72EA", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x976f2B5a5c9aC8007e8c189C4E984dad7bd560a0", amount: "22275", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0x70BD29d01c89686139B55B0bf804FB55d99bc7be", amount: "20250", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },
	{ name: "WM Captial", address: "0xd9C4854Bd8Bd24cbc1c8183Aae883a41eD73E69F", amount: "22500", start: new Date("2024-12-03 10:00:00"), duration: toSecond(15, "Month") },

]

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
	// await createVesting({
	// 	beneficiary: "0x3A3725Bd73df42858b657F0a9BD96e28054885a6",
	// 	token: Tokens.get("NEXG"),
	// 	amount: "1000",
	// 	start: new Date("2024-12-03 10:00:00"),
	// 	duration: toSecond(3, "Month"),
	// 	role: 0,
	// 	releaseInterval: ReleaseInterval.Day,
	// 	cliff: 0,
	// 	cliffReleasePercentage: 0,
	// });
	//await createTestVesting();
	//await getVestingSchedule();
	//createNexgVesting();
	//totalNEXGNeed();
	await sendNEXGToInvestor();
}

async function createNexgVesting() {
	let total = 0;
	for (let index = 0; index < NEXGVestings.length; index++) {
		const v = NEXGVestings[index];
		await createVesting({
			beneficiary: v.address,
			token: Tokens.get("NEXG"),
			amount: v.amount,
			start: v.start,
			duration: v.duration,
			role: 0,
			releaseInterval: ReleaseInterval.Day,
			cliff: 0,
			cliffReleasePercentage: 0,
		});
	}
}

async function sendNEXGToInvestor() {
	const [owner] = await ethers.getSigners();
	const nexg = Tokens.get("NEXG");
	const conToken = await hre.ethers.getContractAt(nexg.contractName, nexg.getAddress(), owner);
	for (let index = 0; index < NEXGVestings.length; index++) {
		const v = NEXGVestings[index];
		process.stdout.write(`send ${v.amount} \tNEXG to\t[${v.name}]${v.address} ....`)
		//do transaction
		const totalAmount = hre.ethers.parseUnits(v.amount, nexg.decimals);
		const tx1 = await conToken.transfer(v.address, totalAmount);
		await waitTransaction(tx1.hash)
		//done
		process.stdout.write(`tx: https://polygonscan.com/tx/${tx1.hash}\n\r`)
	}
}

async function totalNEXGNeed() {
	let total = 0
	for (let index = 0; index < NEXGVestings.length; index++) {
		const v = NEXGVestings[index];
		const a = parseFloat(v.amount)
		total += a;
	}
	console.log("total NEXG need: ", total)
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
	//process.exit(0);
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