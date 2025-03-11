package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 7348 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	//只遍历到倒数第三个元素
	for i := 0; i <= len(nums)-3; i++ {
		n1 := nums[i]
		//避免重复
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - n1
		left := i + 1
		right := len(nums) - 1
		for left < right {
			//需要考虑在遍历过程中的重复问题
			if nums[left]+nums[right] < target {
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//避免重复
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			}
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
