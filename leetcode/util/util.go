package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNode(n []int) *ListNode {
	res := &ListNode{}
	prefix := res
	for _, item := range n {
		prefix.Next = &ListNode{Val: item}
		prefix = prefix.Next
	}
	return res.Next
}
