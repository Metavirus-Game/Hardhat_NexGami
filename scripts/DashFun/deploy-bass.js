const hre = require("hardhat");

const TokenName = "BASS"
const token_address = "0xE71196d768Ebaf3A67cC7528C19e9450a0a04a92";
const ContractFile = `contracts/DashFun/DashFun-${TokenName}.sol:${TokenName}`;

async function main() {
	const [owner] = await hre.ethers.getSigners();
	console.log(owner.address);
	//await deploy();
	await verifyToken();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy ${TokenName}`)
	const c = await hre.ethers.deployContract(ContractFile, [], {
		//gasPrice:"500000000000"
	});
	await c.waitForDeployment();
	console.log(
		`${TokenName} deployed to ${c.target}`
	);
}

async function verifyToken() {
	const [owner] = await hre.ethers.getSigners();
	await verify(token_address, []);
}

async function verify(address, args) {
	console.log("verifying contract...", address);
	const result = await hre.run("verify:verify", {
		address: address,
		constructorArguments: args,
		contract: ContractFile
	})
	console.log(result);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})