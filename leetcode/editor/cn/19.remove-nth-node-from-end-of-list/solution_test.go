package leetcode

import (
	"testing"
)

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 3058 👎 0

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	start := head
	nNode := head
	//向右移动 n 次，间隔为 n
	for i := 0; i < n; i++ {
		nNode = nNode.Next
	}
	prefixNode := res
	for nNode != nil {
		nNode = nNode.Next
		prefixNode = start
		start = start.Next
	}
	prefixNode.Next = start.Next
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {

}
