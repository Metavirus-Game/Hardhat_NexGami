const hre = require("hardhat");

async function main() {
	const [owner] = await ethers.getSigners();
	await deploy()
	//await verify("0xcd7BCaCc38d71ED14C875d3aBFec5a781812551E", [owner.address]);
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy MVT`)
	const myth = await hre.ethers.deployContract("MetaVirusToken", [owner.address], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`MVT deployed to ${myth.target}`
	);
}

async function verify(address, args) {
	console.log("verifying contract...", address);
	const result = await hre.run("verify:verify", {
		address: address,
		constructorArguments: args,
		contract: "contracts/MVT-Token.sol:MetaVirusToken"
	})
	console.log(result);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})