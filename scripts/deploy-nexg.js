const hre = require("hardhat");

async function main() {
	const [owner] = await hre.ethers.getSigners();
	console.log(owner.address);
	// await deploy();
	await verifyNEXG();
}

async function deploy() {
	const [owner] = await hre.ethers.getSigners();
	console.log(owner.address);
	// const nexg = await hre.ethers.deployContract("NexGami", [owner.address], {
	// 	//gasPrice: "100000"
	// });
	const nexg = await hre.ethers.deployContract("NexGami");
	await nexg.waitForDeployment();
	console.log(
		`NexGami deployed to ${nexg.target}`
	);
}

const token_address = "0xaF0dC42725db75AE54f5E8945e71017Bc7ACd27d";

async function verifyNEXG() {
	const [owner] = await hre.ethers.getSigners();
	await verify(token_address, []);
}

async function verify(address, args) {
	console.log("verifying contract...", address);
	const result = await hre.run("verify:verify", {
		address: address,
		constructorArguments: args,
		contract: "contracts/NEXG-Token.sol:NexGami"
	})
	console.log(result);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})