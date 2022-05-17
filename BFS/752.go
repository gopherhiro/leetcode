package BFS

// 752. Open the Lock
// 752. 打开转盘锁
// 思路：BFS
func openLock(deadends []string, target string) int {
	if len(deadends) == 0 || len(target) == 0 {
		return -1
	}
	// 把deadends存入hashtable中，方便后续查找。
	dead := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		dead[v] = true
	}

	// 记录已访问过的密码，防止走回头路。
	visited := make(map[string]bool)
	queue := make([]string, 0)

	step := 0
	start := "0000"

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		rowSize := len(queue)
		for i := 0; i < rowSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 判断密码是否合法，不合法跳过
			if _, ok := dead[current]; ok {
				continue
			}
			// 判断是否找到目标密码，是的话，则返回步数。
			if current == target {
				return step
			}

			// 将节点未遍历的相邻节点加入队列。
			for j := 0; j < 4; j++ {
				up := upOne(current, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}

				down := downOne(current, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	// 如果穷举完所有可能的密码，都没有找到目标密码，则无法解锁。
	return -1
}

func upOne(s string, j int) string {
	runes := []rune(s)
	if s[j] == '9' {
		runes[j] = 48 // 48是'0'的ASCII值
	} else {
		runes[j]++
	}
	return string(runes)
}

func downOne(s string, j int) string {
	runes := []rune(s)
	if s[j] == '0' {
		runes[j] = 57 // 57是'9'的ASCII值
	} else {
		runes[j]--
	}
	return string(runes)
}
