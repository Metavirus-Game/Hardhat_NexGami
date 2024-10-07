const hre = require("hardhat");

const token_amoy = {
	//usdt
	tokenIn: "0xab7c6Db6cfA4D758A8e2e28f372dB43aE316E02A",
	//nexu
	tokenOut: "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3",
	ratio: 1
}

const token_mainnet = {
	//usdt
	tokenIn: "0xc2132D05D31c914a87C6611C10748AEb04B58e8F",
	//nexu
	tokenOut: "0xEb75F3952273B97Aa98fe85a3dD447BC34D4B4De",
	ratio: 1
}

const tokenCfg = {
	polygon_mainnet: token_mainnet,
	polygon_amoy: token_amoy
}

const token = tokenCfg[hre.network.name];

async function main() {
	if (token == null) {
		throw "Token Config Error " + hre.network.name
	}
	await deploy();
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy TokenSwapPermit`)
	const myth = await hre.ethers.deployContract("TokenSwapPermit", [token.tokenIn, token.tokenOut, 1], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`TokenSwapPermit deployed to ${myth.target}`
	);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})