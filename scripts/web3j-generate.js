const fs = require('fs')
const cp = require("child_process");

function help() {
	console.log("\nUsage: node scripts/web3j-generate.js --contract=<Contract Name(no ext)> --package=<Java Package Name>\n")
}

async function main() {
	const args = require('minimist')(process.argv.slice(2))
	if (args["contract"] == null || args["package"] == null) {
		help();
		return;
	}

	const { contract, package } = args;

	//load contract json
	const contractJson = `./artifacts/contracts/${contract}.sol/${contract}.json`;
	if (!fs.existsSync(contractJson)) {
		console.log(`Contract ${contract} Json File Not Found, Please compile the sol file first.`)
		return
	}

	const outDir = `./web3j-out/java/${contract}`
	fs.mkdirSync(outDir, { recursive: true });


	const r = cp.execSync("npx hardhat compile")
	console.log("compiling sol...")
	console.log(r.toString())

	const str = fs.readFileSync(contractJson).toString()
	const json = JSON.parse(str)

	const abi = json.abi;
	const bin = json.bytecode

	const abiFile = `${outDir}/${contract}.abi`
	const binFile = `${outDir}/${contract}.bin`

	fs.writeFileSync(`${abiFile}`, JSON.stringify(abi))
	fs.writeFileSync(`${binFile}`, bin)


	const web3jR = cp.execSync(`web3j generate solidity -a=${abiFile} -b=${binFile} -o=${outDir} -p=${package}`);
	console.log(web3jR.toString());

	console.log("done");
}



main().catch(error => {
	console.error(error);
	process.exitCode = 1;
})