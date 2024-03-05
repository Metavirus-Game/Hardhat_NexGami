const hre = require("hardhat");

async function main() {
	await deploy();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	const myth = await hre.ethers.deployContract("NexGamiUSD", [owner.address], {
		gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`NEXU deployed to ${myth.target}`
	);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})