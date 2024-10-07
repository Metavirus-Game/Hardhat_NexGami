const hre = require("hardhat");

const addr_amoy = {
	//usdt
	tokenIn: "0xab7c6Db6cfA4D758A8e2e28f372dB43aE316E02A",
	tokenInDecimals: 6,
	//nexu
	tokenOut: "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3",
	tokenOutDecimals: 18,
	//
	contractAddr: "0x5BB31759ad675b52e8Cd6A481A096a66f20c040C",
	ratio: 1
}

const addr_mainnet = {
	//usdt
	tokenIn: "0xc2132D05D31c914a87C6611C10748AEb04B58e8F",
	tokenInDecimals: 6,
	//nexu
	tokenOut: "0xEb75F3952273B97Aa98fe85a3dD447BC34D4B4De",
	tokenOutDecimals: 18,
	//
	contractAddr: "0x2815Ec3083f47f073DE9EB5670CfCbF1BD1709dc",
	ratio: 1
}

const addrCfg = {
	polygon_mainnet: addr_mainnet,
	polygon_amoy: addr_amoy
}

const swapAddrs = addrCfg[hre.network.name];

module.exports = { swapAddrs }