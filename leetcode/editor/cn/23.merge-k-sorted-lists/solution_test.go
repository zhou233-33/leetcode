package leetcode

import (
	"fmt"
	"sync"
	"testing"
)

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2971 👎 0

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode, resultChan chan<- *ListNode, wg *sync.WaitGroup) {
	defer wg.Done()
	res := &ListNode{}
	prefix := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prefix.Next = list1
			prefix = prefix.Next
			list1 = list1.Next
		} else {
			prefix.Next = list2
			prefix = prefix.Next
			list2 = list2.Next
		}
	}
	if list1 == nil {
		prefix.Next = list2
	}
	if list2 == nil {
		prefix.Next = list1
	}
	resultChan <- res.Next
}
func mergeKLists(lists []*ListNode) *ListNode {
	var wg sync.WaitGroup
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		newLists := make([]*ListNode, 0)
		resultChan := make(chan *ListNode, len(lists)/2)
		for i := 0; i < len(lists); i = i + 2 {
			//has two
			if i+1 < len(lists) {
				wg.Add(1)
				go func(i int) {
					mergeTwoLists(lists[i], lists[i+1], resultChan, &wg)
				}(i)
			} else {
				//only one left
				newLists = append(newLists, lists[i])
			}
		}
		wg.Wait()
		close(resultChan)
		for item := range resultChan {
			newLists = append(newLists, item)
		}
		lists = newLists
	}
	return lists[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeKSortedLists(t *testing.T) {
	node1 := GenerateListNode([]int{1, 4, 5})
	node2 := GenerateListNode([]int{1, 3, 4})
	node3 := GenerateListNode([]int{2, 6})

	v := mergeKLists([]*ListNode{node1, node2, node3})
	fmt.Println(v)
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
