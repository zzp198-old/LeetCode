package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node1 := list1
	node2 := list1
	for i := 0; i < b; i++ {
		if i < a-1 {
			node1 = node1.Next
		}
		node2 = node2.Next
	}
	node3 := list2
	for node3.Next != nil {
		node3 = node3.Next
	}
	node1.Next = list2
	node3.Next = node2.Next
	return list1
}
