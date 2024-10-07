require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades');
require("hardhat-deploy");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  defaultNetwork: "polygon_amoy",
  solidity: "0.8.23",
  networks: {
    polygon_amoy: {
      url: "https://polygon-amoy.g.alchemy.com/v2/" + process.env.KEY_ALCHEMY_POLYGON_MUMBAI,
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    },
    polygon_mainnet: {
      url: "https://polygon-mainnet.g.alchemy.com/v2/" + process.env.KEY_ALCHEMY_POLYGON_MAINNET,
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    },
    xlayer_test: {
      url: "https://testrpc.xlayer.tech",
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    },
    xlayer: {
      url: "https://rpc.xlayer.tech",
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    },
    BSC_test: {
      url: "https://data-seed-prebsc-1-s1.binance.org:8545",
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    },
    local:{
      url:"http://localhost:30303/",
      accounts: [process.env.NG_Deployer_ETH, process.env.NEXG_Test, process.env.NEXG_PlatformOp]
    }
  },
  etherscan: {
    apiKey: process.env.ETH_SCAN_KEY,
    customChains: [
      {
        network: "xlayer_test",
        chainId: 195,
        urls: {
          apiURL: "https://www.oklink.com/api/explorer/v1/contract/verify/async/api/xlayer_test",
          browserURL: "https://www.oklink.com/xlayer-test"
        }
      },
      {
        network: "xlayer",
        chainId: 196,
        urls: {
          apiURL: "https://www.oklink.com/api/explorer/v1/contract/verify/async/api/xlayer",
          browserURL: "https://www.oklink.com/xlayer"
        }
      }
    ]
  },
  sourcify: {
    enabled: true
  }
};
