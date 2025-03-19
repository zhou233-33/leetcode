package leetcode

import (
	"fmt"
	"runtime"
	"testing"
)

//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
//k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
//,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 3140 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if nums[0] > target && nums[len(nums)-1] < target && nums[0] > nums[len(nums)-1] {
		return -1
	}
	if target == nums[0] {
		return 0
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		midIndex := (left + right) / 2
		midValue := nums[midIndex]
		if midValue == target {
			return midIndex
		}
		if target > nums[0] {
			if midValue < target {
				if midValue >= nums[0] {
					left = midIndex + 1
					continue
				}
				if midValue <= nums[len(nums)-1] {
					right = midIndex - 1
					continue
				}
			} else {
				right = midIndex - 1
			}
			continue
		}
		if target < nums[len(nums)-1] {
			if midValue < target {
				left = midIndex + 1
			} else {
				if midValue >= nums[0] {
					left = midIndex + 1
					continue
				}
				if midValue <= nums[len(nums)-1] {
					right = midIndex - 1
					continue
				}
			}
			continue
		}
	}
	runtime.BlockProfile()
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArray(t *testing.T) {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 4))
}
