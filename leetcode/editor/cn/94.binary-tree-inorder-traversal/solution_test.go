package leetcode

import (
	"fmt"
	"testing"
)

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 2218 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
type myStack []*TreeNode

func (s *myStack) Top() *TreeNode {
	return (*s)[len(*s)-1]
}
func (s *myStack) Push(t *TreeNode) {
	(*s) = append(*s, t)
}
func (s *myStack) Pop() {
	*s = (*s)[:len(*s)-1]
}
func inorderTraversal(root *TreeNode) []int {
	stack := myStack{}
	cur := root
	res := make([]int, 0)
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		res = append(res, stack.Top().Val)
		cur = stack.Top().Right
		stack.Pop()
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 2}
	c := &TreeNode{Val: 3}
	b.Left = c
	a.Right = b
	fmt.Println(inorderTraversal(a))
}
