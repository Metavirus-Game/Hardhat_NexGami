const fs = require("fs")
const hre = require("hardhat");
const ERC20 = require("../ERC20.json");
const { sleep } = require("./include/hre_utils");

const usdc = "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"

const send = async () => {
	const to = "0xd1327a3bd9c4d31a3eddcd826ef24fd2232a7e03";
	const [owner] = await ethers.getSigners();
	const data = {
		to: to,
		value: hre.ethers.parseEther("0.1")
	}
	const tx = await owner.sendTransaction(data)
	console.log(tx);

}

const main = async () => {
	const revokes = readCsv("./腾讯云DMC_数据导出_1730422544954.csv")
	const holder = "0xBB68aFC97Aa8299D0ce29Dee94CE8cEa6D27B529";


	const [owner] = await ethers.getSigners();
	const oPol = await hre.ethers.provider.getBalance(owner.address)
	console.log("Owner is:" + owner.address, "POL=>" + hre.ethers.formatEther(oPol))

	for (const idx in revokes) {
		const revoke = revokes[idx];
		const b = await usdcBalanceOf(revoke.address)
		if (b > 0) {
			let pol = await hre.ethers.provider.getBalance(revoke.address)
			if (pol == 0) {
				const data = {
					to: revoke.address,
					value: hre.ethers.parseEther("0.1")
				}
				console.log("ready to send 0.1 POL to " + revoke.address)
				await sleep(1000)
				const tx = await owner.sendTransaction(data)
				console.log("send 0.1 POL to " + revoke.address, "tx is: ", tx.hash)

				await tx.wait();
				pol = await hre.ethers.provider.getBalance(revoke.address)
			} else if (pol < hre.ethers.parseEther("0.05")) {
				const data = {
					to: revoke.address,
					value: hre.ethers.parseEther("0.05")
				}
				console.log("ready to send 0.05 POL to " + revoke.address)
				await sleep(1000)
				const tx = await owner.sendTransaction(data)
				console.log("send 0.05 POL to " + revoke.address, "tx is: ", tx.hash)

				await tx.wait();
				pol = await hre.ethers.provider.getBalance(revoke.address)
			}
			console.log(revoke.address, " USDC=> " + b, "POL=>" + hre.ethers.formatEther(pol))

			await usdcSend(revoke.getPrivateKey(), holder, hre.ethers.formatUnits(b, 6))

		}
	}

}

async function usdcBalanceOf(address) {
	const [owner] = await ethers.getSigners();
	const provider = hre.ethers.provider;

	const usdcCon = new hre.ethers.Contract(usdc, ERC20.abi, provider);
	const balance = await usdcCon.balanceOf(address);
	return balance
}

/**
 * 
 * @param {string} from sender, pk
 * @param {string} to receiver, address 
 * @param {string} amount usdc amount，单位是1usdc
 */
async function usdcSend(from, to, amount) {
	const provider = hre.ethers.provider;
	const wallet = new hre.ethers.Wallet(from, provider);
	const usdcCon = new hre.ethers.Contract(usdc, ERC20.abi, wallet);
	const amt = hre.ethers.parseUnits(amount, 6)
	const tx = await usdcCon.transfer(to, amt, {
		maxPriorityFeePerGas: hre.ethers.parseUnits("25", "gwei"), // 优先费设置为 25 Gwei
		maxFeePerGas: hre.ethers.parseUnits("80", "gwei") // 最大 Gas 费设置为 50 Gwei
	})
	console.log("transfer " + amount + " usdc to " + to, "tx hash: ", tx.hash)
	await tx.wait();
}



class RevokeUser {
	/**
	 * @type {number}
	 */
	id;
	/**
	 * @type {string}
	 */
	address;
	/**
	 * @type {string}
	 */
	mnemonic;
	/**
	 * @type {string}
	 */
	privateKey;

	constructor(id, address, mnemonic, privateKey) {
		this.id = id;
		this.address = address;
		this.mnemonic = mnemonic;
		this.privateKey = privateKey;
	}

	getPrivateKey() {
		if (this.privateKey.length < 64) {
			this.privateKey = this.privateKey.padStart(64, "0")
		}
		return this.privateKey
	}
}

/**
 * 
 * @param {string} csvFile 
 * @returns {RevokeUser[]}
 */
const readCsv = (csvFile) => {
	const exist = fs.existsSync(csvFile);
	/**
	 * @type {RevokeUser[]}
	 */
	const revokes = []
	const buf = fs.readFileSync(csvFile)
	const content = buf.toString("utf8");
	const data = content.split("\n");
	for (const idx in data) {
		const d = data[idx].replace("\r", "").split(",")
		//const [id, address, mnemonic, privateKey] = d;
		const ru = new RevokeUser(...d);
		let found = false;
		revokes.push(ru)

	}
	return revokes;
}


main().catch((e) => {
	console.error(e);
})