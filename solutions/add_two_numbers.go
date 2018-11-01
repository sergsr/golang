package solutions

func addTwoNumbers(l1 *given.ListNode, l2 *given.ListNode) *given.ListNode {
	preHead := &given.ListNode{}
	current := preHead
	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		sum := carry + l1.Val + l2.Val
		carry = sum / 10
		current.Next = &given.ListNode{Val: sum % 10}
		current = current.Next
	}

	longest := l1
	if longest == nil {
		longest = l2
	}

	for ; longest != nil; longest = longest.Next {
		sum := carry + longest.Val
		carry = sum / 10
		current.Next = &given.ListNode{Val: sum % 10}
		current = current.Next
	}

	for carry > 0 {
		current.Next = &given.ListNode{Val: carry % 10}
		current = current.Next
		carry /= 10
	}

	return preHead.Next
}
