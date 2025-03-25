package sort

// reversePairs 计算逆序对数量的主函数
func reversePairs(record []int) int {
	if len(record) < 2 {
		return 0
	}
	// 创建临时数组，避免在递归中反复分配
	temp := make([]int, len(record))
	return mergeSort(record, 0, len(record)-1, temp)
}

// mergeSort 递归分治并统计逆序对
func mergeSort(record []int, left, right int, temp []int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	// 递归计算左半部分和右半部分的逆序对
	count := mergeSort(record, left, mid, temp) + mergeSort(record, mid+1, right, temp)
	// 合并并统计跨左右部分的逆序对
	count += merge(record, left, mid, right, temp)
	return count
}

// merge 合并两个有序子数组并统计逆序对
func merge(record []int, left, mid, right int, temp []int) int {
	// 初始化指针
	i := left    // 左半部分起始
	j := mid + 1 // 右半部分起始
	k := left    // 临时数组指针
	count := 0   // 逆序对计数器

	// 合并两个子数组
	for i <= mid && j <= right {
		if record[i] <= record[j] {
			temp[k] = record[i]
			i++
		} else {
			// record[i] > record[j]，产生逆序对
			temp[k] = record[j]
			count += mid - i + 1 // 左半部分 [i...mid] 都大于 record[j]
			j++
		}
		k++
	}

	// 处理剩余元素
	for i <= mid {
		temp[k] = record[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = record[j]
		j++
		k++
	}

	// 将临时数组复制回原数组
	for p := left; p <= right; p++ {
		record[p] = temp[p]
	}

	return count
}
