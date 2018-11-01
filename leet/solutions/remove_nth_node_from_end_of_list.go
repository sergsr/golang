package solutions

import (
	"github.com/sergsr/golang/leet/given"
)

func removeNthFromEnd(head *given.ListNode, n int) *given.ListNode {
	preHead := &given.ListNode{Next: head}
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
