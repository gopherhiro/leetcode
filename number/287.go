package number

// 287. Find the Duplicate Number
// 287. 寻找重复数
// 思路：XOR异或
// x ^ 0 = x
// x ^ x = 0
// x ^ y = y ^ x
func findDuplicateXOR(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	for i := 0; i < len(nums); i++ {
		ans ^= i
	}

	return ans
}
