package leetcode

import (
	"fmt"
	"testing"
)

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 2942 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 找到大于 target 的最小数， 小于 target 的最大数
func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	resLeft := -1
	resRight := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			if nums[mid] == target {
				resRight = mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			if nums[mid] == target {
				resLeft = mid
			}
			right = mid - 1
		}
	}
	return []int{resLeft, resRight}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
