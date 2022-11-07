package linkedlist

// 147. Insertion Sort List
// 147. 对链表进行插入排序
// 思路：插入排序
// time O(n^2) space O(1)
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	helper := &ListNode{} //new starter of the sorted list
	pre := helper         //insert node between pre and pre.next
	cur := head           //the node will be inserted

	for cur != nil {
		next := cur.Next

		//find the right place to insert
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}

		//insert between pre and pre.next
		cur.Next = pre.Next
		pre.Next = cur
		pre = helper
		cur = next
	}
	return helper.Next
}
