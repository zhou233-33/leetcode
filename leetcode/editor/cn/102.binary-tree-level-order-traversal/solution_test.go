package leetcode

import (
	"testing"
)

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 2109 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a binary tree node.
//* type TreeNode struct {
//*     Val int
//*     Left *TreeNode
//*     Right *TreeNode
//* }
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	layer := make([]*TreeNode, 0)
	layer = append(layer, root)
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	for len(layer) > 0 {
		curLen := len(layer)
		tmpRes := make([]int, 0)
		for i := 0; i < curLen; i++ {
			item := layer[i]
			tmpRes = append(tmpRes, item.Val)
			if item.Left != nil {
				layer = append(layer, item.Left)
			}
			if item.Right != nil {
				layer = append(layer, item.Right)
			}
		}
		res = append(res, tmpRes)
		layer = layer[curLen:]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {

}
