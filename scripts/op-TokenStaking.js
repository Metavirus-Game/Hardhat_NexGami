const hre = require("hardhat");
const { verify, waitTransaction } = require("./include/hre_utils");
const { Tokens } = require("./include/tokens")

const _nftAddress = {
	polygon_amoy: "0x14788ed721B80cDD8c32cC69C6c46432a2Ef5568",
	BSC_test: "0x62E5F868b0b0B51Cb056a3e7B60241494D985A21"
}

const nftAddress = () => {
	return _nftAddress[hre.network.name]
}

const stakingAddr = "0x522a63afc828F96731616208966B415E653d5E36"

async function main() {
	//①先DeployNFT
	//await deployNFT();

	//②在Deploy Staking合约
	const addr = await deployStaking();

	//③对Staking合约进行设置
	await configStaking(addr);

	//④在nft中给staking合约增加Minter角色
	await addMinterRole(addr);

	await getStakingTokenInfo(addr);
}

/**
 * 在nft合约中给staking合约分配MINTER角色
 * @param {staking的合约地址} contractAddress 
 */
async function addMinterRole(contractAddress) {
	const [owner] = await ethers.getSigners();
	const contract = await hre.ethers.getContractAt("StakeRewardTest", nftAddress(), owner);
	const tx = await contract.grantRole(contract.MINTER_ROLE(), contractAddress)
	await waitTransaction(tx.hash)
	console.log(`grant MINTER_ROLE to [${contractAddress}] tx[${tx.hash}]`)
}

async function getStakingTokenInfo(address) {
	const [owner] = await ethers.getSigners();
	//set other tokens and rewardNft
	const contract = await hre.ethers.getContractAt("MultiTokenStaking", address, owner);
	const tx = await contract.getTokenInfo()
	console.group(`Staking[${address}] Token Info`);
	console.log(tx)
	console.groupEnd();
}

async function deployNFT() {
	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy Test NFT`)
	const nft = await hre.ethers.deployContract("StakeRewardTest", [owner.address, owner.address], {
		//gasPrice:"500000000000"
	});
	await nft.waitForDeployment();
	console.log(
		`Test NFT deployed to ${nft.target}`
	);
}

async function configStaking(address) {
	const [owner] = await ethers.getSigners();
	//set other tokens and rewardNft
	const contract = await hre.ethers.getContractAt("MultiTokenStaking", address, owner);

	const otherTokens = [
		Tokens.get("MVT").getAddress()
	];
	const ratios = [
		hre.ethers.parseUnits("10", 18) //10倍基础token
	];
	const rewardNft = nftAddress() //奖励的nft


	const tx = await contract.configureStaking(otherTokens, ratios, rewardNft);
	await waitTransaction(tx.hash);
}

async function deployStaking() {
	const baseToken = Tokens.get("NEXG").getAddress();
	const rewardToken = Tokens.get("NEXU").getAddress();
	const annualInterestRate = 500; //年利率5%
	const nftThreshold = hre.ethers.parseUnits("1000", 18);  //至少质押100个token才能获得nft
	const minStakeDuration = 60 * 60;// 最少质押1小时
	const minStakeAmount = hre.ethers.parseUnits("100", 18); //最少质押token数量

	const [owner] = await ethers.getSigners();
	console.log(`Owner[${owner.address}] Deploy Multi Staking`)
	const staking = await hre.ethers.deployContract("MultiTokenStaking",
		[
			baseToken,
			rewardToken,
			annualInterestRate,
			nftThreshold,
			minStakeDuration,
			minStakeAmount
		]);
	await staking.waitForDeployment();
	console.log(
		`Multi Staking deployed to ${staking.target}`
	);

	return staking.target;
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})