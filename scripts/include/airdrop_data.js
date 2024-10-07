
class AirDropData {
	/**
	 * @type {string}
	 * 
	 */
	title;
	/**
	 * @type {string}接收方钱包地址
	 */
	address;
	/**
	 * @type {number} token数量，此处为最大精度单位，例如(ether)
	 */
	tokenAmount;

	/**
	 * 出现次数，多次只发一次
	 * @bug bug_0603 多次出现是因为参与了多次活动，不能只发一次
	 */
	count;

	/**
	 * 由于统计多个地址时出现错误，此处记录>1次出现地址的应发token总数，进行补发
	 */
	remainingAmount;

	/**
	 * token的实际数量，最小精度单位
	 */
	get amount() {
		return ethers.parseEther(this.tokenAmount.toString()).toString()
	}

	constructor(title, address, tokenAmount) {
		this.title = title;
		this.count = 1;
		this.address = address;
		this.tokenAmount = parseFloat(tokenAmount);
		this.remainingAmount = 0;
	}
}

module.exports = AirDropData