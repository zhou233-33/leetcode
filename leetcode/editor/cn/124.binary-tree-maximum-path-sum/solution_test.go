package leetcode

import (
	"math"
	"testing"
)

//二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定
//经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//
//
// 提示：
//
//
// 树中节点数目范围是 [1, 3 * 10⁴]
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 2379 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	calculateMaxPath(root, &res)
	return res
}

// calculateMaxPath return maxPath end with root
func calculateMaxPath(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftVal := calculateMaxPath(root.Left, res)
	rightVal := calculateMaxPath(root.Right, res)
	maxPathValue := max(leftVal+rightVal+root.Val, leftVal+root.Val, rightVal+root.Val, root.Val)
	*res = max(*res, maxPathValue)
	return max(leftVal+root.Val, rightVal+root.Val, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeMaximumPathSum(t *testing.T) {

}
