require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades');
require("hardhat-deploy");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  defaultNetwork: "polygon_mumbai",
  solidity: "0.8.23",
  networks: {
    polygon_mumbai: {
      url: "https://polygon-amoy.g.alchemy.com/v2/"  + process.env.KEY_ALCHEMY_POLYGON_MUMBAI,
      accounts: [process.env.NG_Deployer_ETH]
    },
    polygon_mainnet: {
      url: "https://polygon-mainnet.g.alchemy.com/v2/" + process.env.KEY_ALCHEMY_POLYGON_MAINNET,
      accounts: [process.env.NG_Deployer_ETH]
    }
  },
  etherscan: {
    apiKey: process.env.ETH_SCAN_KEY
  },
  sourcify: {
    enabled: true
  }
};
