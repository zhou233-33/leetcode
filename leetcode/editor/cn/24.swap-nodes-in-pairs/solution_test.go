package leetcode

import (
	"fmt"
	"testing"
)

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
// Related Topics 递归 链表 👍 2374 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{}
	prefix := res
	left := head
	if left == nil {
		return head
	}
	right := head.Next
	if right == nil {
		return head
	}
	for left != nil && right != nil {
		//swap
		left.Next = right.Next
		right.Next = left
		prefix.Next = right
		//record prefix element
		prefix = left
		if left.Next != nil {
			left = left.Next
		} else {
			break
		}
		if left.Next != nil {
			right = left.Next
		} else {
			break
		}
	}
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {
	n := []int{1}
	head := GenerateListNode(n)
	res := swapPairs(head)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
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
