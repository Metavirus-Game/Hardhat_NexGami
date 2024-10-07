const hre = require("hardhat");
const { verify } = require("./include/hre_utils");

async function main() {
	await deploy();
	// await verifyNEXU();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy Nexu`)
	const myth = await hre.ethers.deployContract("NexGamiUSD", [owner.address], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`NEXU deployed to ${myth.target}`
	);
}

const token_address = "0xEb75F3952273B97Aa98fe85a3dD447BC34D4B4De";

async function verifyNEXU() {
	const [owner] = await hre.ethers.getSigners();
	await verify(token_address, "contracts/NEXU.sol:NexGamiUSD", [owner.address]);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})