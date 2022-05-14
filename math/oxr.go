package math

// 思路：XOR异或
// x ^ 0 = x
// X ^ x = 0
// x ^ y = y ^ x
func findDuplicateR(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	for i := 0; i < len(nums); i++ {
		ans ^= i
	}

	return ans
}

