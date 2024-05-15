const { ethers, upgrades } = require("hardhat");

async function main() {
	await deploy();
	//await upgrade();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(owner.address);

	const nexg = await ethers.getContractFactory("NexGami");
	console.log("deploying NexGami....")

	const m = await upgrades.deployProxy(nexg, [owner.address], {
		kind: "uups",
		initializer: "initialize",
	});
	await m.waitForDeployment();
	//await nexg.waitForDeployment();
	console.log(
		`NexGami deployed to ${m.target}`
	);
}

const PROXY = "0x0Df7F618718A65Fed01c9f0B70E37fFD6cA66D14"
async function upgrade() {
	const mV2 = await ethers.getContractFactory("NexGami");
	var m = await upgrades.upgradeProxy(PROXY, mV2);
	console.log("NexGami upgraded successfully", m.target);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})