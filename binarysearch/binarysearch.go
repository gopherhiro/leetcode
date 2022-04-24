package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 5, 6}
	res := searchInsert(nums, 7)
	fmt.Println(res)
}

// 704. Binary Search
// 思路：使用二分搜索框架即可。
// time O(logN), space O(1)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 中间值小于target，表示target的值在右边。
			// 故而收缩左边，即去右边找。
			left = mid + 1
		} else if nums[mid] > target {
			// 中间值大于target，表示target的值在左边。
			// 故而收缩右边，即去左边找。
			right = mid - 1
		} else if nums[mid] == target {
			return mid // 找到，直接返回。
		}
	}
	return -1
}

// 34. Find First and Last Position of Element in Sorted Array
// 34. 在排序数组中查找元素的第一个和最后一个位置
// 思路：使用二分搜索框架
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := leftBound(nums, target)
	right := rightBound(nums, target)
	return []int{left, right}
}

// 搜索左侧边界的二分搜索
func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不返回，收缩右边界，从而锁定左边界。
			right = mid - 1
		}
	}
	// 检查越界情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 搜索右侧边界的二分搜索
func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不返回，收缩左边界，从而锁定右边界。
			left = mid + 1
		}
	}
	// 检查越界情况
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 35. 搜索插入位置
// 35. Search Insert Position
// 思路：搜索左侧边界的二分搜索。
// 搜索左侧边界的二分搜索，当目标元素 target 不存在数组 nums 中时，
// 搜索左侧边界的二分搜索的返回值可以做以下几种解读：
// 1、返回的这个值是 nums 中大于等于 target 的最小元素索引。
// 2、返回的这个值是 target 应该插入在 nums 中的索引位置。
// 3、返回的这个值是 nums 中小于 target 的元素个数。
// time O(logN) space O(1)
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不返回，收缩右边界，从而锁定左边界。
			right = mid - 1
		}
	}
	return left
}

// 392. 判断子序列
// 392. Is Subsequence
// 思路：双指针 time (N + M),space O(1)
// 初始化两个指针 i 和 j，分别指向 s 和 t 的初始位置；
// 每次贪心地匹配，匹配成功则 i 和 j 同时右移，匹配 s 的下一个位置；
// 匹配失败，则 j 右移，i 不变，尝试用 t 的下一个字符匹配 s；
// 最终如果 i 移动到 s 的末尾，就说明 s 是 t 的子序列。
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// 278. 第一个错误的版本
// 思路：二分左边界搜索
// 注意性质：当一个版本为正确版本，则该版本之前的所有版本均为正确版本；
// 当一个版本为错误版本，则该版本之后的所有版本均为错误版本。
func firstBadVersion(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return true
}

// 374. Guess Number Higher or Lower
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		answer := guess(mid)
		if answer == 1 {
			left = mid + 1
		} else if answer == -1 {
			right = mid - 1
		} else if answer == 0 {
			return mid
		}
	}
	return 0
}

func guess(num int) int {
	return 0
}

// 172. 阶乘后的零
// 思路：数学原理，计算乘法因子中5的个数。
func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}

// 793. 阶乘函数后 K 个零
// 思路：搜索有多少个 n 满足 trailingZeroes(n) == K，
// 其实就是在问，满足条件的 n 最小是多少，最大是多少，最大值和最小值一减，就可以算出来有多少个 n 满足条件了。
// 那就是 二分查找 中「搜索左侧边界」和「搜索右侧边界」呀
// 观察题目给出的数据取值范围，n 可以在区间 [0, LONG_MAX] 中取值，
// 寻找满足 trailingZeroes(n) == K 的左侧边界和右侧边界， 左侧边界和右侧边界之差 + 1即是答案。
func preimageSizeFZF(k int) int {
	min := leftBorder(k)
	max := rightBorder(k)
	return max - min + 1
}

func leftBorder(target int) int {
	left, right := 0, math.MaxInt64
	for left <= right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid - 1
		} else if trailingZeroes(mid) == target {
			right = mid - 1
		}
	}
	return left
}

func rightBorder(target int) int {
	left, right := 0, math.MaxInt64
	for left <= right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid - 1
		} else if trailingZeroes(mid) == target {
			left = mid + 1
		}
	}
	return right
}

// 875. Koko Eating Bananas
// 875. 爱吃香蕉的珂珂
// 思路：二叉搜索之左侧边界搜索。
func minEatingSpeed(piles []int, limitHour int) int {
	if len(piles) == 0 || limitHour == 0 {
		return 0
	}
	// left,right 表示吃香蕉的速度范围
	left, right := 1, getMax(piles)
	for left <= right {
		mid := left + (right-left)/2
		// 以 mid 速度能否在 limitHour 小时内吃完香蕉
		if canFinish(piles, mid, limitHour) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 以  speed 速度，能否在 limitHour 小时内吃完所有香蕉
func canFinish(piles []int, speed, limitHour int) bool {
	sumHour := 0
	for _, number := range piles {
		sumHour += timeOf(number, speed)
	}
	return sumHour <= limitHour
}

// 以 speed 速度吃 number 个香蕉，需要的时间
func timeOf(number, speed int) (needTime int) {
	needTime = number / speed
	// 如果一小时内吃不完，则只能下一个小时再吃。
	// 所以，一次吃不完的，还需要加1个小时来吃完。
	if number%speed > 0 {
		needTime += 1
	}
	return
}

// 计算数组的最大值
func getMax(nums []int) int {
	max := 0
	for _, v := range nums {
		max = maxvalue(max, v)
	}
	return max
}

func maxvalue(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 1011. Capacity To Ship Packages Within D Days
// 1011. 在 D 天内送达包裹的能力
// 思路：二分搜索之左侧边界，寻找运载能力最小值。
// time O(N*logM), N 是 weights中的元素个数， M是其所有元素之和。
func shipWithinDays(weights []int, limitedDays int) int {
	if len(weights) == 0 || limitedDays == 0 {
		return 0
	}
	// 船的最小载重应该是weights数组中元素的最大值，因为每次至少得装一件货物走。
	// 最大载重显然就是weights数组所有元素之和，也就是一次把所有货物都装走。
	left, right := getMax(weights), getSum(weights)
	for left <= right {
		mid := left + (right-left)/2
		if canFinishShip(weights, mid, limitedDays) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left
}

// 以 capacity 的运载能力，能否在限定的天数内完成运输
func canFinishShip(weights []int, capacity, days int) bool {
	// needDays 初始值最小为1(题目要求)。
	sum, needDays := 0, 1
	for _, weight := range weights {
		if sum+weight > capacity {
			needDays++
			sum = 0
		}
		sum += weight
	}
	return needDays <= days
}

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 410. Split Array Largest Sum
// 410. 分割数组的最大值
// 思路：二分搜索之左侧边界，寻找子数组和最大的最小值。
// time O(N*logM)，N 是nums元素的个数，M是其所有元素之和。
// 难点在于实际问题如何转换为二分搜索问题。
func splitArray(nums []int, m int) int {
	if len(nums) == 0 || m == 0 {
		return 0
	}
	// 子数组至少包含一个元素，至多包含整个数组
	// 题目要求：子数组和最大，所以：
	// 「最大」子数组和的取值范围就是闭区间[max(nums), sum(nums)]
	left, right := getMax(nums), getSum(nums)
	for left <= right {
		mid := left + (right-left)/2
		count := splitCount(nums, mid)
		if count < m {
			// 最大子数组和上限高了，减小一些，所以，收缩右边界。
			right = mid - 1
		} else if count > m {
			// 最大子数组和上限低了，增加一些，所以，收缩左边界。
			left = mid + 1
		} else if count == m {
			// 收缩右边界，从而锁定左边界。
			right = mid - 1
		}
	}
	return left
}

// 若限制最大子数组和为 max，
// 计算 nums 至少可以被分割成几个子数组
func splitCount(nums []int, max int) int {
	sum, count := 0, 1
	for _, v := range nums {
		if sum+v > max {
			count++
			sum = 0
		}
		sum += v
	}
	return count
}

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

// 162. Find Peak Element
// 162. 寻找峰值
// 思路：二分搜索
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	// A peak element is an element
	// that is strictly greater than its neighbors.
	// the last element, which does not meet the requirements, is ignored.
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// in a descending sequence of numbers. or a local falling slope
			// reduce the search space to the left of mid(including itself)
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 852. Peak Index in a Mountain Array
// 与 162 相同。
func peakIndexInMountainArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// in a descending sequence of numbers. or a local falling slope
			// reduce the search space to the left of mid(including itself)
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 33. Search in Rotated Sorted Array
// 33. 搜索旋转排序数组
// 思路：「旋转数组中找目标值」 转化成「有序数组中找目标值」
// 使用基础二分搜索即可。
// time(logN)
func searchRotate(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		// 先确定 target 在左半段还是右半段
		// 更改mid的值，构造有序搜索范围
		// 从而将「旋转数组中找目标值」 转化成了「有序数组中找目标值」
		if target >= nums[0] {
			// target 在左半段
			// 如果mid在右半段的话，则将mid改为 max
			if nums[mid] < nums[0] {
				nums[mid] = math.MaxInt32
			}

		} else {
			// target 在右半段
			// 如果mid在左半段的话，则将mid改为 min
			if nums[mid] >= nums[0] {
				nums[mid] = math.MinInt32
			}
		}
		// 经过以上操作，已经转换为有序数组找目标值
		// 故而使用基础二分搜索解决问题即可。
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 思路：将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。
// 无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。
// 以此类推。
func searchRotateN(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		// 确认有序部分，在有序中进行二分搜索
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

// 69. x 的平方根
// 69. Sqrt(x)
// only the integer part of the result is returned.
// 思路：在 0 ~ x 范围内进行二分搜索即可。
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 2, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else if mid > x/mid {
			right = mid - 1
		}
	}
	return right
}

// 287. Find the Duplicate Number
// 287. 寻找重复数
// 思路：n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内。
// 索引是有序的，所以，可以使用二分查找。
func findDuplicateQ(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 1, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1

		// 统计小于等于mid的个数
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		// 如果count <= mid,表示左侧数据不存在重复的数
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}

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
