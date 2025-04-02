package leetcode

import (
	"fmt"
	"testing"
)

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2920 👎 0

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

func isSymmetric(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root.Left, root.Right)
	for len(q) > 0 {
		l := q[0]
		r := q[1]
		q = q[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSymmetricTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 2}
	//left2 := &TreeNode{Val: 3}
	right1 := &TreeNode{Val: 3}
	//right2 := &TreeNode{Val: 3}
	root.Left = left1
	root.Right = right1
	//left1.Right = left2
	//right1.Right = right2
	fmt.Println(isSymmetric(root))
}
