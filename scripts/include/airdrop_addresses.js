const hre = require("hardhat");
const ERC20 = require("../../ERC20.json");

const polygon_amoy = {
	ad_address: "0xe31D3f3432CB3715DF2ceA6Db91F0E2b320015aa",
	token_address: "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3",
	cost_per_address: 0.0005
}

const polygon_mainnet = {
	ad_address: "0x60797243339a08958E71ed78A3fDB8f170560B9b",
	token_address: "0xaF0dC42725db75AE54f5E8945e71017Bc7ACd27d",
	// token_address: "0x08c11a08c9F0994beCF64369CBe866d88c029545", //测试用的nexu
	cost_per_address: 0.004
}

const _address = {
	polygon_amoy,
	polygon_mainnet
};

const { ad_address, token_address, cost_per_address } = _address[hre.network.name];

async function tokenBalanceOf(walletAddress) {
	const provider = hre.ethers.provider;
	const token = new hre.ethers.Contract(token_address, ERC20.abi, provider);
	const balance = await token.balanceOf(walletAddress);
	return balance;
}

module.exports = { ad_address, token_address, cost_per_address, tokenBalanceOf }
