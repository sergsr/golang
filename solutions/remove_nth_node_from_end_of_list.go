package solutions

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{Next: head}
	left := preHead
	right := preHead
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}
