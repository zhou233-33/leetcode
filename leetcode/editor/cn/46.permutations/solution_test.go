package leetcode

import (
	"fmt"
	"testing"
)

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 3068 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	traverse(nums, tmp, &res)
	return res
}
func traverse(nums []int, tmp []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, tmp)
		return
	}
	for index, item := range nums {
		nextTmp := tmp
		nextTmp = append(nextTmp, item)
		copyNum := make([]int, len(nums))
		copy(copyNum, nums)
		copyNum = deleteItem(copyNum, index)
		traverse(copyNum, nextTmp, res)
	}
}

func deleteItem(nums []int, index int) []int {
	nums[index] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutations(t *testing.T) {
	fmt.Println(permute([]int{1, 2}))
}
