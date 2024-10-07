const fs = require("fs")

/**
 * 随机xp分数，每1个xp分数=6.94NEXG，共20万NEXG
 */

const xpPoints = {
	tele_ann: { point: 3, max: 1, percentage: 50 },
	tele_group: { point: 3, max: 1, percentage: 50 },
	discord: { point: 3, max: 1, percentage: 50 },
	x_follow: { point: 5, max: 1, percentage: 90 },
	x_repost: { point: 5, max: 1, percentage: 60 },
	x_retweet: { point: 2, max: 1, percentage: 40 },
	x_like: { point: 2, max: 1, percentage: 70 },
	x_quote: { point: 5, max: 1, percentage: 30 },
	refer: { point: 10, max: 10, percentage: 30 },
	refer_extra: { point: 10, min: 30, max: 2000, percentage: 5 }//10%几率附加一个大量refer的用户
}

const NEXG_MAX = 10000000;
const XP_VALUE = 6.94;

const main = async () => {
	let totalNexg = 0;

	const users = [];
	while (totalNexg < NEXG_MAX) {
		const u = randomUser();
		// if (totalNexg + parseFloat(u.nexg) > NEXG_MAX) {
		// 	break;
		// }
		users.push(u);
		totalNexg += parseFloat(u.nexg);
	}

	let content = "";

	users.forEach(u => {
		content += `${u.point}\n`
	})

	console.log(`total ${users.length} users ${totalNexg} NEXG`)
	fs.writeFileSync("./fake-users.csv", content);
}

const randomUser = () => {
	const user = {
		xpItems: {},
		point: 0,
		get nexg() {
			return (parseInt(this.point) * XP_VALUE).toFixed(2);
		}
	}
	for (const key in xpPoints) {
		const item = xpPoints[key];
		const r = Math.random() * 100;
		if (r > item.percentage) {
			continue;
		}
		const min = item.min ?? 1;
		const max = item.max;
		let count = min;
		if (min != max) {
			count = parseInt(Math.random() * (max - min))
		}
		const p = item.point * count;
		user.point += p;
		user.xpItems[key] = {
			point: item.point,
			count: count,
			score: p,
		}
		// console.log(`用户获得 ${key} ${r} ${r < item.percentage} -- ${p}分`)
	}
	return user;
}

main().catch(e => {
	console.error(e);
	process.exit(1);
})