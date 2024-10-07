const hre = require("hardhat");

class TokenDef {
	/**
	 * @type {string}
	 */
	name;

	/**
	 * @type {{networkName:address}}
	 */
	contract;


	/**
	 * Hardhat中的合约名字
	 * @type {string}
	 */
	contractName;
	/**
	 * @type {number}
	 */
	decimals;

	/**
	 * @type {()=>string}
	 */
	getAddress;
}


const USDT = {
	name: "USDT",
	//token的合约地址，network name -> 合约地址
	contract: {
		polygon_mainnet: "0xc2132D05D31c914a87C6611C10748AEb04B58e8F", //Polygon
		polygon_amoy: "0xab7c6Db6cfA4D758A8e2e28f372dB43aE316E02A" //Polygon-amoy
	},
	decimals: 6,
}

const NEXU = {
	name: "NEXU",
	//token的合约地址，chainId -> 合约地址
	contract: {
		polygon_mainnet: "0xEb75F3952273B97Aa98fe85a3dD447BC34D4B4De",
		polygon_amoy: "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3",
		local: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		BSC_test: "0x0F81EF40e9666D6f802c72C1B6dA7Cba1938E82e"
	},
	contractName: "NexGamiUSD",
	decimals: 18,
}

const NEXTEST = {
	name: "NEXTEST",
	//token的合约地址，chainId -> 合约地址
	contract: {
		polygon_mainnet: "0x0Db6Ed389824a10Bc79b2A5296b81b118f07B9d5",
		polygon_amoy: "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3",
		local: "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	},
	contractName: "NexGamiTest",
	decimals: 18,
}

const NEXG = {
	name: "NEXG",
	//token的合约地址，chainId -> 合约地址
	contract: {
		polygon_mainnet: "0xaF0dC42725db75AE54f5E8945e71017Bc7ACd27d",
		polygon_amoy: "0x60797243339a08958E71ed78A3fDB8f170560B9b",
		local: "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
		BSC_test: "0x60797243339a08958E71ed78A3fDB8f170560B9b"
	},
	contractName: "NexGamiUUPS",
	decimals: 18,
}

const MVT = {
	name: "MVT",
	//token的合约地址，chainId -> 合约地址
	contract: {
		polygon_mainnet: "0xcd7BCaCc38d71ED14C875d3aBFec5a781812551E",
		polygon_amoy: "0x36Ec469229Ee2AdF88FF3CC66d9f66cbD5b7Cb69",
		local: "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
		BSC_test: "0x8D7dfD82B324C2d3C1a1B7123F885c8F45716745"
	},
	contractName: "MetaVirusToken",
	decimals: 18,
}

const TicketAndStaking = {
	name: "TicketAndStaking",
	//token的合约地址，chainId -> 合约地址
	contract: {
		polygon_mainnet: "0x514E7999875D6acd8b92DE79Fc99C070eB907c97",
		polygon_amoy: "0xF48A375AF138b3a25674a0f073224Ba8a0ac25d0",
		local: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	},
	decimals: 18,
}

const MetaVirusVIPStake = {
	name: "MetaVirusVIPStake",
	contract: {
		polygon_mainnet: "0x61346562d7E58C15C95d091B077f7d83aFbA8c7E",
		polygon_amoy: "0x0cBE365c0C90C1C59D096ef1186875dA01e5C1D2",
	}
}

const MVP = {
	name: "MetaVirusPet NFT",
	contract: {
		polygon_amoy: "0x42B2C52B5D175b97c0B51AB5aA7DF1b4ebabdECd",
		polygon_mainnet: "0x7A390343006F894A01Da02942228EBDB2465E92F"
	}
}

/**
 * 
 * @param {string} name 
 * @returns {TokenDef}
 */
const get = (name) => {
	const t = Tokens[name];
	if (t == null) return null;
	if (t.getAddress == null) {
		t.getAddress = () => {
			return t.contract[hre.network.name];
		}
	}
	return t;
}

const Tokens = {
	USDT, NEXG, NEXU, MVT, NEXTEST, MVP, TicketAndStaking, MetaVirusVIPStake,
	get
}

module.exports = { Tokens, TokenDef }