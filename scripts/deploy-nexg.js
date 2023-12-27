const hre = require("hardhat");

async function main() {
	await deploy();
}

async function deploy() {
	const nexg = await hre.ethers.deployContract("NexGami");
	await nexg.waitForDeployment();
	console.log(
		`NexGami deployed to ${nexg.target}`
	);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})