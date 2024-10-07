const hre = require("hardhat");
const fs = require("fs")
const ERC20 = require("../ERC20.json");
const colors = require("colors-console")
const AirDropData = require("./include/airdrop_data")

const { ad_address, token_address, cost_per_address, tokenBalanceOf } = require("./include/airdrop_addresses")

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
	return airDrops;
}


const csvInvestors = { in: "./TGE Send Investors.csv", out: "./token-sent-investors.csv" };
const csvUsers = { in: "./TGE Send Users.csv", out: "./token-sent-users.csv" }
const csvUserInner = { in: "./TGE Send Users Inner.csv", out: "./token-sent-inner.csv" }
const csvExchanges = { in: "./TGE Send Exchanges.csv", out: "./token-sent-exchanges.csv" };

const currentCsvFile = csvUsers;
const main = async () => {
	const airDrops = readCsv(currentCsvFile.in)
	let r = "Title,Address,Token Sent,Token Balance\n";

	for (const idx in airDrops) {
		const ad = airDrops[idx];
		if (!ethers.isAddress(ad.address)) {
			continue;
		}
		const balance = await tokenBalanceOf(ad.address);
		const log = `${ad.title},${ad.address},${ad.tokenAmount + ad.remainingAmount},${hre.ethers.formatEther(balance)}`
		r += log + "\n";
		process.stdout.write(`${idx}/${airDrops.length}			\r`)
	}

	console.log(r);
	fs.writeFileSync(currentCsvFile.out, r);
}

main().catch(e => {
	console.error(e);
	process.exitCode = 1
})