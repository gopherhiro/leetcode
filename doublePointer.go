package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	s := "accaabaa"
	res := longestPalindrome(s)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 88. Merge Sorted Array
// 88. 合并两个有序数组
// 思路：two pointer
// time O(m+n) space O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		// 判断是否到达 num1 末尾
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		// 判断是否到达 num2 末尾
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

// 5. Longest Palindromic Substring
// 5. 最长回文子串
// 思路：中心拓展法
// time O(N^2) space O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		// 处理回文串长度为奇数的情况：
		// 以 s[i] 为中心的最长回文子串
		si := palindrome(s, i, i)
		// 处理回文串长度为偶数的情况：
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		st := palindrome(s, i, i+1)
		// 比较选出最长的回文子串
		if len(si) > len(ans) {
			ans = si
		}
		if len(st) > len(ans) {
			ans = st
		}
	}
	return ans
}

// 在 s 中寻找以 s[left] 和 s[right] 为中心的最长回文串
// 1、当 left == right 时，就是在寻找长度为奇数的回文串。
// 2、当 left != right 时，就是在寻找长度为偶数的回文串。
func palindrome(s string, left, right int) string {
	// 从中心向外扩展
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// 9. Palindrome Number
// 9. 回文数
// 思路：int to string, and two pointer
// time O(n) space O(1)
func isPalindromeINT(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return checkPalindrome(s, 0, len(s)-1)
}

// 680. Valid Palindrome II
// 680. 验证回文字符串 Ⅱ
// 思路：
// 1、先验证 s 是否是回文串
// 2、如果 s 是回文串，则return true
// 3、如果 s 不是回文串，则把 s 分成 2 份。
// 即：[left..right-1] 和 [left+1..right] 分别验证回文性。
// time O(N), space O(1)
func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// 去掉一个字符情况分为：[left..right-1] 和 [left+1..right] 区域
			// 分别 check 以上两种情况即可。
			return checkPalindrome(s, left, right-1) || checkPalindrome(s, left+1, right)
		}
		left++
		right--
	}
	return true
}

// 验证回文性
func checkPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 125. Valid Palindrome
// 125. 验证回文串
// 思路：two pointer
// 1、去掉所有非数字、字母字符。
// 2、统一转为 lower case or upper case
// 3、判断回文
// time O(N) space O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 去掉所有非数字、字母字符。
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = re.ReplaceAllString(s, "")
	// to lower case
	s = strings.ToLower(s)
	// two pointer 判断是否回文
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 125. Valid Palindrome
// 125. 验证回文串
// 思路：在原字符串上使用 two pointer
// time O(N) space O(1)
func isPalindromeO(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)

	left, right := 0, len(s)-1
	for left < right {
		// find left alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// find right alphanumeric characters
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 判断是否是字符或数字
func isAlphanumeric(ch byte) bool {
	return ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch >= '0' && ch <= '9'
}

// 283. Move Zeroes
// 283. 移动零
// 思路：two pointer
// time O(N) space O(1)
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		// 使用 slow 指针标记值为 0 的元素位置，
		// 使用 fast 指针找到值不为 0 的元素的位置
		// 交换两个即可。
		if nums[fast] != 0 {
			// [0..slow]都不为0
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

// 备注：这道题有点不好理解，可以忽略。
// 80. Remove Duplicates from Sorted Array II
// 80. 删除有序数组中的重复项 II
// 思路：two pointer
// 原地修改：
// 1、slow 指针指向当前即将放置元素的位置。
// 2、fast 指针向后遍历所有元素。
// time O(N) space O(1)
func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		// 因为数组是有序的，如果 num[fast] == num[slow-2]
		// 则出现 num[fast] == num[slow-2] == num[slow-1]
		// 存在连续的多于2个的重复元素。
		if nums[fast] == nums[slow-2] {
			// 所以，当 num[fast] == num[slow-2]时
			// slow 不动，标记当前需要删除元素的位置
			// fast 则继续后移，找寻新元素。
			fast++
		} else {
			// 使用新元素替换 slow 位置的元素
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	// slow 即是去重后的数组长度
	return slow
}

// 27. Remove Element
// 27. 移除元素
// 思路：two pointer
// time O(N) space O(1)
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != val {
			// 维护 [0..slow-1] 不重复
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// slow 即是去重后的数组长度
	return slow
}

// 82. Remove Duplicates from Sorted List II
// 82. 删除排序链表中的重复元素 II
// 思路：two pointer
// 1、找到重复区域，2、跳过重复区域。
// time O(N), space O(1)
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	virHead := &ListNode{}
	virHead.Next = head

	slow, fast := virHead, virHead.Next
	for fast != nil {

		// fast 指针先走，找到重复的区域。
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}

		// 如果 slow 没有紧跟着 fast，表示中间有重复元素，需要跳跃。
		if slow.Next != fast {
			// 跳过重复区域元素，即删除重复元素。
			slow.Next = fast.Next
		} else {
			// slow 跟上 fast
			slow = slow.Next
		}

		// 继续遍历
		fast = fast.Next
	}
	return virHead.Next
}

// 83. Remove Duplicates from Sorted List
// 83. 删除排序链表中的重复元素
// 思路：双指针
// time O(N), space O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

// 26. Remove Duplicates from Sorted Array
// 26. 删除有序数组中的重复项
// 思路：双指针
// time O(N) space O(1)
func removeDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			// 维护 nums[0..slow] 无重复
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 557. Reverse Words in a String III
// 557. 反转字符串中的单词 III
// 思路：以空格分割句子成单词数组，
// 遍历单词数组，reverse每一个单词。
// time O(N), space O(N)
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Split(s, " ")
	newWords := make([]string, 0)
	for _, word := range words {
		t := []byte(word)
		reverseString(t)
		newWords = append(newWords, string(t))
	}
	return strings.Join(newWords, " ")
}

// 345. Reverse Vowels of a String
// 345. 反转字符串中的元音字母
// 思路：双指针
// time O(N) space O(1)
func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	t := []byte(s)

	left, right := 0, len(t)-1
	for left < right {
		// 找到左边的元音字母
		for left < right && !isVowel(t[left]) {
			left++
		}
		// 找到右边的元音字母
		for left < right && !isVowel(t[right]) {
			right--
		}
		// left, right 都找到，则进行交换
		// 交换之后继续遍历
		t[left], t[right] = t[right], t[left]
		left++
		right--
	}

	return string(t)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

// 541. Reverse String II
// 541. 反转字符串 II
// 思路：双指针
// time (N) space O(N)
func reverseStr(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return s
	}

	t := []byte(s)

	for i := 0; i < len(t); i += 2 * k {
		// 计算交换范围
		// right = i+k-1，表示右边界范围
		// right 与 len(t) - 1 取 min，
		// 可保证不足 k 个时，将剩余字符全部反转。
		left, right := i, min(i+k-1, len(t)-1)
		for left < right {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}
	}

	return string(t)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Reverse Kth String
// 反转前k个字符
// 思路：左右指针，交换k-1个元素即可。
// time O(N) space (1)
func reverseKthStr(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return s
	}
	ns := []byte(s)
	left, right := 0, k-1
	for left < right {
		ns[left], ns[right] = ns[right], ns[left]
		left++
		right--
	}
	return string(ns)
}

// 344. Reverse String
// 344. 反转字符串
// 思路：左右指针，交换元素即可。
// time O(N), space O(1)
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 反转数组
// 思路：左右指针
// time O(N) space O(1)
func reverseArray(nums []int) {
	if len(nums) == 0 {
		return
	}
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// 167. Two Sum II - Input Array Is Sorted
// 167. 两数之和 II - 输入有序数组
// 思路：双指针（左右指针）。
// time O(N) space O(1)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 203. 移除链表元素
// 思路：遍历链表，查找删除即可。
// 创建虚拟头节点，方便删除。
// time O(N) space O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 使用虚拟节点，方便删除
	virHead := &ListNode{}
	virHead.Next = head

	prev := virHead

	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return virHead.Next
}

// 2095. Delete the Middle Node of a Linked List
// 2095. 删除链表的中间节点
// 思路：快慢指针
// time O(N) space O(1)
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 使用快慢指针，找出中间节点和其前驱节点
	slow, fast := head, head

	// 记录中间节点的前驱节点
	prev := &ListNode{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	// slow 是中间节点
	// prev 是中间节点的前一个节点
	// 删除中间节点
	prev.Next = slow.Next
	return head
}

// 面试题 02.02. 返回倒数第 k 个节点
// 思路：快慢指针
// time O(N) space O(1)
func kthToLast(head *ListNode, k int) int {
	if head == nil || k == 0 {
		return -1
	}
	// 让快指针先都 k 步
	// 然后快慢指针一起走
	// 快指针到达链表末尾时，慢指针就是倒数第k个位置。
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Val
}

// 876. Middle of the Linked List
// 876. 链表的中间结点
// 思路：快慢指针
// time O(N), space O(1)
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 当fast到达链表末尾时，slow即是中点。
	return slow
}

// 142. Linked List Cycle II
// 142. 环形链表 II
// 思路：快慢指针
// time O(N) space O(1)
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 如果不存在环，则返回 null
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 如果存在环
	// 则让其中一个指针从头节点重新出发
	slow = head
	for slow != fast {
		// 两指针以相同的速度前进
		slow = slow.Next
		fast = fast.Next
	}
	// 两指针相遇的节点，即是环的起点
	return slow
}

// 141. Linked List Cycle
// 141. 环形链表
// 思路：快慢指针
// time O(N), space O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 快慢指针指向头节点
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		// 快指针每次前进两步
		fast = fast.Next.Next

		// 慢指针每次前进一步
		slow = slow.Next

		// 快指针、慢指针相遇，链表存在环
		if fast == slow {
			return true
		}
	}
	return false
}

// 创建链表
// 尾插法
// time O(N) space O(1)
func createListFromTail(nums []int) *ListNode {
	var head, tail, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, value := range nums {
		// 新建节点，并指向尾节点
		node := &ListNode{Val: value}
		node.Next = tail
		curr.Next = node
		curr = node
	}
	return head.Next
}

// 创建链表
// 头插法
// time O(N), space O(1)
func createListFromHead(nums []int) *ListNode {
	var prev, curr *ListNode
	for _, value := range nums {
		node := &ListNode{}
		node.Val = value
		node.Next = prev
		curr = node
		prev = curr
	}
	return curr
}

func printLinkedList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d->", head.Val)
		} else {
			fmt.Printf("%d", head.Val)
		}
		head = head.Next
	}
	fmt.Println()
}
