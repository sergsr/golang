package solutions

import (
	"container/heap"

	"github.com/sergsr/golang/leet/given"
)

type listHeap []*given.ListNode

func (h listHeap) Len() int { return len(h) }

func (h listHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h listHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *listHeap) Push(x interface{}) {
	*h = append(*h, x.(*given.ListNode))
}

func (h *listHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*given.ListNode) *given.ListNode {
	preHead := &given.ListNode{}
	current := preHead
	var frontier listHeap
	for _, x := range lists {
		if x != nil {
			frontier = append(frontier, x)
		}
	}
	heap.Init(&frontier)

	for frontier.Len() > 0 {
		next := heap.Pop(&frontier).(*given.ListNode)
		current.Next = next
		current = current.Next
		if next.Next != nil {
			heap.Push(&frontier, next.Next)
		}
	}

	return preHead.Next
}
