const hre = require("hardhat");
const { Tokens } = require("./include/tokens");
const { sleep, waitTransaction } = require("./include/hre_utils");

async function main() {
	//await deploy();
	await setEventDetails();
	// await getClaimedUsers();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy TokenSwap`)

	const testMode = false

	const ticketToken = Tokens.get("NEXG");
	console.log(ticketToken)
	const stakeToken = Tokens.get("NEXU");
	// const stakeToken = Tokens.get("NEXTEST");
	const rewardToken = Tokens.get("MVT");

	let ticketAmount = hre.ethers.parseEther("100")
	const minStakeAmount = hre.ethers.parseEther("1")
	const maxStakeAmount = hre.ethers.parseEther("10000")

	const totalStakeLimit = hre.ethers.parseEther("10000000")
	const tgePercentage = 10 //tge释放10%代币
	const rewardRatio = 20 * 100; //1 stakeToken = 20 rewardToken，合约中按照百分比计算，所以20倍要x100
	let startTime = parseInt(new Date("2024-06-15 20:00:00").getTime() / 1000)
	let duration = 3600 * 24
	if (testMode) {
		//测试时间改这里
		startTime = parseInt(new Date("2024-06-15 11:00:00").getTime() / 1000)
		duration = 3600 * 24
	}

	//const refundDuration = 86400

	process.stdout.write("Deploying TicketAndStaking Contract\n")
	if (testMode) {
		//测试模式不收nexg
		ticketAmount = hre.ethers.parseEther("0")
		process.stdout.write("TEST MODE!!!!!!!!\n")
		process.stdout.write("TEST MODE!!!!!!!!\n")
		process.stdout.write("TEST MODE!!!!!!!!\n")
		process.stdout.write("TEST MODE!!!!!!!!\n")
	}
	process.stdout.write(`Ticket:\t${ticketToken.name} -- ${ticketToken.getAddress()} ${hre.ethers.formatEther(ticketAmount)}\n`)
	process.stdout.write(`Stake:\t${stakeToken.name} -- ${stakeToken.getAddress()} ${hre.ethers.formatEther(minStakeAmount)} to ${hre.ethers.formatEther(maxStakeAmount)}\n`)
	process.stdout.write(`Reward:\t${rewardToken.name} -- ${rewardToken.getAddress()} Ratio: ${rewardRatio}\n`)
	process.stdout.write(`Total Raised:\t ${hre.ethers.formatEther(totalStakeLimit)}\n`)

	process.stdout.write(`Reward Ratio:\t${(rewardRatio / 100.0).toFixed(2)}倍\n`)
	process.stdout.write(`TGE Percentage:\t${tgePercentage}%\n`)

	process.stdout.write(`Start Time:\t${new Date(startTime * 1000).toLocaleDateString()} ${new Date(startTime * 1000).toLocaleTimeString()}\n`)
	process.stdout.write(`End Time:\t${new Date((startTime + duration) * 1000).toLocaleDateString()} ${new Date((startTime + duration) * 1000).toLocaleTimeString()}\n`)
	//process.stdout.write(`Refund Ended:\t${new Date((startTime + duration + refundDuration) * 1000).toLocaleDateString()} ${new Date((startTime + duration + refundDuration) * 1000).toLocaleTimeString()}\n`)

	let ans = ""

	if (testMode) {
		process.stdout.write(`Current Is TESTMODE, Continue[y/n]?`)
	} else {
		process.stdout.write('Continue[y/n]?')
	}
	process.stdin.on('data', (input) => {
		const a = input.toString().trim();
		if (['Y', 'y', 'YES', 'yes'].indexOf(a) > -1) {
			ans = "y"
		} else {
			ans = "n"
		}
	})

	process.stdout.write("\n")

	while (ans == "") {
		await sleep(1000)
	}

	if (ans != 'y') {
		process.exit(0)
	}

	const c = await hre.ethers.deployContract("TicketAndStaking",
		[
			{
				ticketToken: ticketToken.getAddress(),
				ticketAmount: ticketAmount,
				stakeToken: stakeToken.getAddress(),
				minStakeAmount: minStakeAmount,
				maxStakeAmount: maxStakeAmount,
				totalStakeLimit: totalStakeLimit,
				rewardToken: rewardToken.getAddress(),
				tgePercentage,
				rewardTokenRatio: rewardRatio,
				startTime,
				duration,
			}
		]
	)
	console.log("Deploying contract.....")
	await c.waitForDeployment();
	console.log(
		`TicketAndStaking deployed to ${c.target}`
	);
	process.exit(0)
}

const contractAddress = "0x0Da7859f9760956c7aA852d3f60728e11CF11Cba"

//2024-06-16
//合约有bug，之前refundDruation参数没有彻底去掉，claim逻辑不应该有期限
async function setEventDetails() {
	const [owner] = await ethers.getSigners();
	let startTime = parseInt(new Date("2024-06-15 20:00:00").getTime() / 1000)
	let duration = 3600 * 24
	let refundDuration = 3600 * 24 * 14


	const con = await hre.ethers.getContractAt("TicketAndStaking", contractAddress, owner);
	const tx = await con.setEventDetails(startTime, duration, refundDuration)
	console.log(tx);
	const r1 = await waitTransaction(tx.hash);
	if (!r1) {
		console.error("setEventDetails failed");
		process.exit(1);
	}
	console.log("done")
}

async function getContractBalances() {
	const [owner] = await ethers.getSigners();
	const con = await hre.ethers.getContractAt("TicketAndStaking", contractAddress, owner);
	const tx = await con.getContractBalances()
	console.log(tx);
}

async function getClaimedUsers(){
	const [owner] = await ethers.getSigners();
	const con = await hre.ethers.getContractAt("TicketAndStaking", contractAddress, owner);
	const tx = await con.getClaimedUsers()
	console.log(tx);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})