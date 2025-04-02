package leetcode

import (
	"fmt"
	"testing"
)

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
// 返回其根节点。
//
// 示例 1:
//
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 示例 2:
//
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
//
// 提示:
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2504 👎 0
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for index, val := range inorder {
		inorderMap[val] = index
	}
	return GenerateNode(preorder, inorder, inorderMap, 0, 0, len(inorder)-1)
}

func GenerateNode(preorder []int, inorder []int, inorderMap map[int]int, preIndex, midL, midR int) *TreeNode {
	if preIndex >= len(preorder) || midL > midR {
		return nil
	}
	res := &TreeNode{}
	for preIndex < len(preorder) {
		preFirst := preorder[preIndex]
		inorderIndex := inorderMap[preFirst]
		if inorderIndex > midR || inorderIndex < midL {
			preIndex++
			continue
		}
		res.Val = preFirst
		res.Left = GenerateNode(preorder, inorder, inorderMap, preIndex+1, midL, inorderIndex-1)
		res.Right = GenerateNode(preorder, inorder, inorderMap, preIndex+1, inorderIndex+1, midR)
		break
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	res := buildTree(preorder, inorder)
	fmt.Println(res)
}
