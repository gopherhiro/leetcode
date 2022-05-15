package twopointer

import (
	"fmt"
	"leetcode/pkg"
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
		left, right := i, pkg.Min(i+k-1, len(t)-1)
		for left < right {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}
	}

	return string(t)
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
