const hre = require("hardhat");
const ERC20 = require("../ERC20.json");
const nexu_address = "0x5E086c0E963eA27027E2Ecd573f0994Aa22Cf9d3"
async function main() {
	await balance();
}

async function balance() {
	const [owner] = await ethers.getSigners();
	const provider = hre.ethers.provider;
	const usdcCon = new hre.ethers.Contract(nexu_address, ERC20.abi, provider);
	const balance = await usdcCon.balanceOf(owner.address);
	console.log("Balance ", balance);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})