package leetcode

import (
	"fmt"
	"testing"
)

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 2503 👎 0

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{}
	prefix := res
	//TODO we need to think special case: head is nil, k=1
	if k == 1 || head == nil {
		return head
	}
	left := head
	right := head
OuterLoop:
	for left != nil && right != nil {
		//move right, k = move k-1 times
		for i := 0; i < k-1; i++ {
			right = right.Next
			if right == nil {
				break OuterLoop
			}
		}
		front := left
		back := front.Next
		rightNext := right.Next
		//reverse items between left and right
		for front != right {
			backNext := back.Next
			back.Next = front
			front = back
			back = backNext
		}
		// process the left and right node
		prefix.Next = right
		left.Next = rightNext
		//move on
		prefix = left
		left = rightNext
		right = left
	}
	return res.Next
}

// leetcode submit region end(Prohibit modification and deletion)
func GenerateListNode(n []int) *ListNode {
	res := &ListNode{}
	prefix := res
	for _, item := range n {
		prefix.Next = &ListNode{Val: item}
		prefix = prefix.Next
	}
	return res.Next
}

func TestReverseNodesInKGroup(t *testing.T) {
	n := []int{1, 2, 3, 4, 5}
	head := GenerateListNode(n)
	res := reverseKGroup(head, 3)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}
