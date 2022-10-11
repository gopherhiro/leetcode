package number

// 1518. 换酒问题
// 思路：模拟：减法
// time O(N) space O(1)
func numWaterBottles(numBottles int, numExchange int) int {
	answer, emptyBottles := numBottles, numBottles
	for emptyBottles >= numExchange {
		emptyBottles = emptyBottles - numExchange // 每进行一次兑换
		answer++                                  // 就有一瓶酒喝
		emptyBottles++                            // 就产生一个空瓶
	}
	return answer
}

// 1518. 换酒问题
// 思路：模拟：除法
// time O(logN) space O(1)
func numWaterBottles2(bottle int, exchange int) int {
	answer := bottle
	// 当且仅当空瓶子个数 bottle 大于等于兑换个数 exchange 时，可以继续喝到酒
	for bottle >= exchange {
		wine := bottle / exchange   // 兑换后得到酒的个数
		remain := bottle % exchange // 兑换后剩余的酒瓶子个数
		answer += wine              // 可以喝到的酒增加
		bottle = wine + remain      // 剩余空瓶子个数
	}
	return answer
}
