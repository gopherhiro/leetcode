package array

// 75. Sort Colors
// 75. 颜色分类
// 思路：三指针
// time O(n) space O(1)
func sortColors(nums []int) {
	// [red] [white] [blue]
	// we think the white element is already in correct place.
	// we iterate the array nums:
	// if nums[white] == 0, we swap with red pointer, and red++, white++
	// if nums[white] == 1, we don't have to swap.
	// if nums[white] == 2, we swap with blue pointer, blue--
	red, white, blue := 0, 0, len(nums)-1
	for white <= blue {
		if nums[white] == 0 {
			nums[red], nums[white] = nums[white], nums[red]
			red++
			white++
			continue
		}
		if nums[white] == 1 {
			white++
			continue
		}
		if nums[white] == 2 {
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}

// 75. Sort Colors
// 75. 颜色分类
// 思路：桶排序
// time O(n) space O(n)
func sortColors2(nums []int) {
	// 1.使用3个桶分别存储 0,1,2 数值
	// 2.收集桶中的元素，覆盖 nums 中的数值即可
	buckets := make([][]int, 3)
	for _, num := range nums {
		buckets[num] = append(buckets[num], num)
	}
	tmp := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		tmp = append(tmp, bucket...)
	}
	copy(nums, tmp)
}
