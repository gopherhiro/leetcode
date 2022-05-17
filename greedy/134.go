package greedy

// 134. Gas Station
// 134. 加油站
// 策略：数学图像法
// 思路：思路：剩余油量最低点作为起点，能够保证油量累加和一直大于等于 0。
// time O(N) space O(1)
func canCompleteCircuitN(gas []int, cost []int) int {
	n := len(gas)
	start := 0
	sum, minSum := 0, 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			// 选择油量最低点作为起点
			// 最有可能保证环游一周
			start = i + 1
			minSum = sum
		}
	}
	if sum < 0 {
		// 总增加油量小于总消耗油量，无解
		return -1
	}
	if start == n {
		// 环形数组
		return 0
	}
	return start
}

// 134. Gas Station
// 134. 加油站
// 策略：贪心
// time O(N) space O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		// 总增加油量小于总消耗油量，无解
		return -1
	}

	start := 0
	remain := 0
	for i := 0; i < n; i++ {
		remain += gas[i] - cost[i]
		if remain < 0 {
			start = i + 1 // 更新起点
			remain = 0
		}
	}

	// 环形数组
	if start == n {
		return 0
	}
	return start
}
