const hre = require("hardhat");
const fs = require("fs")
const ERC20 = require("../ERC20.json");
const colors = require("colors-console")

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

const COST_PER_ADDRESS = cost_per_address;

const SEND_STEP = 250;

const TOKEN_DECIMALS = 18

const { ethers } = hre;

/**
 * 读取csv文件，csv文件从excel直接导出
 * 要求第一行是标题行
 * 第二行开始是数据
 * 第一列是名称，第二列是地址，第三列式发送金额
 * @param {string} csvFile csv文件
 * @returns {AirDropData[]}
 */
const readCsv = (csvFile) => {
	const exist = fs.existsSync(csvFile);
	const airDrops = []
	const buf = fs.readFileSync(csvFile)
	const content = buf.toString("utf8");
	const data = content.split("\n");
	for (const idx in data) {
		if (idx == 0) continue; //ignore title row
		const d = data[idx].replace("\r", "").split(",")
		const [title, address, tokenAmount] = d;
		const airDrop = new AirDropData(...d);
		let found = false;
		for (const key in airDrops) {
			const adItem = airDrops[key];
			if (adItem.address == airDrop.address) {
				adItem.count++;
				adItem.remainingAmount += airDrop.tokenAmount;
				found = true;
			}
		}
		if (!found) {
			airDrops.push(airDrop);
		}
	}
	airDrops.forEach(ad => {
		if (ad.count > 1) {
			console.log(`---${ad.address} 重复参与活动 , 共${ad.count}次, 首发${ad.tokenAmount}, 应补发${ad.remainingAmount}`)
		}
	})
	return airDrops;
}

async function tokenBalanceOf(walletAddress) {
	const provider = hre.ethers.provider;
	const usdcCon = new hre.ethers.Contract(token_address, ERC20.abi, provider);
	const balance = await usdcCon.balanceOf(walletAddress);
	return balance;
}

/**
 * 
 * @param {AirDropData[]} airDrops 
 * @param {number} totalAmount 要发送的token总数量
 */
const doAirDrop = async (airDrops, totalAmount) => {
	console.log("Processing AirDrop...")
	const [owner] = await ethers.getSigners();
	const maticBalance = await ethers.provider.getBalance(owner);
	const mbFloat = ethers.formatUnits(maticBalance, "ether");
	const tokenAllowance = await allowance_token();
	const allowFloat = ethers.formatUnits(tokenAllowance, "ether");
	const tokenBalance = await tokenBalanceOf(owner.address);
	const tokenEther = parseFloat(ethers.formatEther(tokenBalance));

	if (tokenEther < totalAmount) {
		console.log(colors(['red', 'bright'], `insufficient Token, sending ${totalAmount} Token, current balance is ${tokenEther}, ${totalAmount - tokenEther} Token needed`))
		process.exit(1);
	}

	const estimatedCost = COST_PER_ADDRESS * airDrops.length;

	if (estimatedCost > mbFloat) {
		console.log(`insufficient Matic, need ${estimatedCost} Matic at least`)
		process.exit(1);
	}

	if (allowFloat < totalAmount) {
		console.log(`not enough allowance of token, requesting approve [${ethers.parseEther(totalAmount.toString())}] token...`);
		const tx = await approve_token(ethers.parseEther(totalAmount.toString()));
		console.log(`approve request sent, waiting for confirm the tx ${tx}`)

		let t = 0;

		while (true) {
			const r = await ethers.provider.getTransactionReceipt(tx);
			process.stdout.write(".");
			if (r == null) {
				await sleep(1000);
				t += 1000;
				if (t > 5 * 60 * 1000) {
					//超时5分钟
					process.stdout.write("Timeout\r\n");
					process.exit(1);
				}
			} else {
				process.stdout.write("Done");
				break;
			}
		}
		console.log("");
	}

	const ad = await hre.ethers.getContractAt("AirDropper", ad_address, owner);

	const destWallet = [];
	const value = [];
	for (const idx in airDrops) {
		const ad = airDrops[idx];
		if (ad.dropped) {
			continue;
		}
		destWallet.push(ad.address);
		value.push(ad.amount);
	}
	if (airDrops.length <= SEND_STEP) {
		const tx = await ad.multisend(token_address, destWallet, value);
		console.log("multi send " + destWallet.length + " wallets");
		console.log("pending tx:", tx.hash);
	} else {
		const step = SEND_STEP;
		for (let idx = 0; idx < destWallet.length; idx += step) {
			const from = idx;
			const to = Math.min(idx + step, destWallet.length);
			const wallets = destWallet.slice(from, to);
			const values = value.slice(from, to);
			console.log("multi send from " + from + " to " + to + " Length[" + wallets.length + "]");
			const tx = await ad.multisend(token_address, wallets, values);
			console.log("pending tx:", tx.hash);
		}
	}
	console.log(colors('green', "Done"))
	process.exit(0);
}

async function approve_token(amount) {
	const [owner] = await ethers.getSigners();
	const provider = ethers.provider;
	const usdcCon = new hre.ethers.Contract(token_address, ERC20.abi, owner);
	const tx = await usdcCon.approve(ad_address, amount);
	return tx.hash;
}

async function allowance_token() {
	const [owner] = await ethers.getSigners();
	const provider = hre.ethers.provider;
	const usdcCon = new hre.ethers.Contract(token_address, ERC20.abi, owner);
	const allow = await usdcCon.allowance(owner.address, ad_address);
	return allow
}

const sleep = async (ms) => {
	return new Promise((r) => {
		setTimeout(() => {
			r();
		}, ms);
	})
}

const csvInvestors = "./TGE Send Investors.csv";
const csvUsers = "./TGE Send Users.csv"
const csvUserInner = "./TGE Send Users Inner.csv"

const currentCsvFile = csvUsers;

/**
 * 设置为true，只发送bug0603中出现错误的用户
 */
const send_bug_0603 = false;

const main = async () => {
	let airDrops = readCsv(currentCsvFile);
	let totalAmount = 0;

	if (send_bug_0603) {
		const nd = [];
		for (const idx in airDrops) {
			const ad = airDrops[idx];
			//bug_0603 只处理count>1的情况
			if (ad.count <= 1) continue;
			console.log(`resend to address\t${ad.address} ${ad.remainingAmount}\tNEXG`)
			if (!ethers.isAddress(ad.address)) {
				console.log(`airdrop data in line ${(parseInt(idx) + 2)} : address[${ad.address}] is not a ether address, removed`)
				ad.dropped = true;
				continue;
			}
			nd.push(new AirDropData(ad.title, ad.address, ad.remainingAmount));
			totalAmount += ad.tokenAmount;
		}
		airDrops = nd;
	} else {
		for (const idx in airDrops) {
			const ad = airDrops[idx];
			if (!ethers.isAddress(ad.address)) {
				console.log(`airdrop data in line ${(parseInt(idx) + 2)} : address[${ad.address}] is not a ether address, removed`)
				ad.dropped = true;
				continue;
			}
			totalAmount += ad.tokenAmount;
		}
	}


	console.log(``)
	if (send_bug_0603) {
		console.log(colors(['red', 'right'], "sending tokens to bug users!!!!"))
	}
	console.log(`current network [${colors(['blue', 'bright'], hre.network.name)}]`);
	console.log(`current address `, _address[hre.network.name]);

	const [owner] = await ethers.getSigners();
	const maticBalance = await ethers.provider.getBalance(owner);

	console.log(`OP Address ${colors('blue', owner.address)}, Matic Balance ${colors('blue', ethers.formatUnits(maticBalance, "ether"))}`)
	process.stdout.write(`total ${colors('green', airDrops.length)} wallet addresses, send ${colors('green', totalAmount)} NEXG, Continue[y/n]?`)

	process.stdin.on('data', (input) => {
		const a = input.toString().trim();
		if (['Y', 'y', 'YES', 'yes'].indexOf(a) > -1) {
			doAirDrop(airDrops, totalAmount);
		} else {
			process.exit(0);
		}
	})
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1;
})