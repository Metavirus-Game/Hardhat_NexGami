const hre = require("hardhat");

async function main() {
    const [owner] = await hre.ethers.getSigners();
    const balance = await hre.ethers.provider.getBalance(owner.address);
    console.log("%s (%f OKX)", owner.address, hre.ethers.formatEther(balance));
    console.log(0.0001 * Math.pow(10,9))
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});