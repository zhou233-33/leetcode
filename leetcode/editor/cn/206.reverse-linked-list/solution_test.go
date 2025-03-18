package leetcode

import (
	"fmt"
	"testing"
)

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 3831 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode1 struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var front *ListNode
	back := head
	for back != nil {
		backNext := back.Next
		back.Next = front
		front = back
		back = backNext
	}
	return front
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedList(t *testing.T) {
	list := GenerateListNode([]int{})
	res := reverseList(list)
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
