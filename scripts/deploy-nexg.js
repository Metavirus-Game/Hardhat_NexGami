const hre = require("hardhat");

async function main() {
	await deploy();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy nexg`)
	const nexg = await hre.ethers.deployContract("NexGami", [], {
		//gasPrice:"500000000000"
	});
	await nexg.waitForDeployment();
	console.log(
		`NEXG deployed to ${nexg.target}`
	);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})