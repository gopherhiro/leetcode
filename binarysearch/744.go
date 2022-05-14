package binarysearch

// 744. 寻找比目标字母大的最小字母
// 744. Find Smallest Letter Greater Than Target
// 思路：二分搜索之右边界搜索
func nextGreatestLetter(letters []byte, target byte) byte {
	if len(letters) == 0 {
		return 0
	}

	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] < target {
			left = mid + 1
		} else if letters[mid] > target {
			right = mid - 1
		} else if letters[mid] == target {
			left = mid + 1
		}
	}
	// 注意环绕case
	if right+1 > len(letters)-1 {
		return letters[0]
	}
	return letters[right+1]
}
