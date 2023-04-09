package offer

// 剑指 Offer 06. 从尾到头打印链表
// 思路：辅助栈
// time O(n) space O(n)
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	answer := make([]int, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		answer = append(answer, stack[i])
	}
	return answer
}
