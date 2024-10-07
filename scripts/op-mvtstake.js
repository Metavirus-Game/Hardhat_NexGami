/**
 * mtvstake是metavirus内部测试时使用的合约
 * 用户用户质押指定数量的mvt，获取对应等级的vip
 * 纯内部用
 */

const { Tokens } = require("./include/tokens");
const hre = require("hardhat");
const { verify, waitTransaction } = require("./include/hre_utils");

async function main() {
	// await deployMVP()
	await deploy();
}

async function deployMVP() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy MetaVirusPet NFT`)

	const minterAddr = Tokens.get("MetaVirusVIPStake").getAddress();

	const myth = await hre.ethers.deployContract("MetaVirusPet", [owner.address, minterAddr], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`MetaVirusPet deployed to ${myth.target}`
	);
}

async function deploy() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy MetaVirusVIPStake`)

	const MVT = Tokens.get("MVT");
	const rewardRate = 500; //万分之500
	const duration = 86400 * 3;//三天

	const MVP = Tokens.get("MVP");//奖励的NFT

	const myth = await hre.ethers.deployContract("MetaVirusVIPStake", [MVT.getAddress(), rewardRate, MVP.getAddress(), duration], {
		//gasPrice:"500000000000"
	});
	await myth.waitForDeployment();
	console.log(
		`MetaVirusVIPStake deployed to ${myth.target}`
	);

	addMinterRole(myth.target);
}

/**
 * 在nft合约中给staking合约分配MINTER角色
 * @param {contract的合约地址} contractAddress 
 */
async function addMinterRole(contractAddress) {
	const MVP = Tokens.get("MVP");//奖励的NFT

	const [owner] = await ethers.getSigners();
	const contract = await hre.ethers.getContractAt("MetaVirusPet", MVP.getAddress(), owner);
	const tx = await contract.grantRole(contract.MINTER_ROLE(), contractAddress)
	await waitTransaction(tx.hash)
	console.log(`grant MINTER_ROLE to [${contractAddress}] tx[${tx.hash}]`)
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})