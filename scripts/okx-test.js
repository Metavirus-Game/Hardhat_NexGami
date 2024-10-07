const hre = require("hardhat");

async function main() {
    const [owner] = await hre.ethers.getSigners();

    tx = await hre.ethers.provider.getTransaction("0xcc8a5c9971d7e96bf51b5e2fb0511fcfe6d96e63fd7308a0a81fabff0fab653b")
    tx.gasPrice=0
    console.log(tx);
    hre.ethers.provider.call(tx);

}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});