const hre = require("hardhat");

async function main() {
	await deploy();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	const myth = await hre.ethers.deployContract("USDT", [owner.address]);
	await myth.waitForDeployment();
	console.log(
		`USDT deployed to ${myth.target}`
	);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})