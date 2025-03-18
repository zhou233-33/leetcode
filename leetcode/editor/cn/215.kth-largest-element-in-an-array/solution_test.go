package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2669 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}
func quickSort(nums []int, left, right, k int) {
	pivot := rand.Intn(right-left+1) + left
	//left elem is bigger than pivot
	pivotNum := nums[pivot]
	nums[pivot], nums[right] = nums[right], nums[pivot]
	index := left
	for i := left; i < right; i++ {
		if nums[i] > pivotNum {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index], nums[right] = nums[right], nums[index]
	if index == k-1 {
		return
	}
	if index > k-1 {
		quickSort(nums, left, index-1, k)
	} else {
		quickSort(nums, index+1, right, k)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAnArray(t *testing.T) {
	n := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	res := findKthLargest(n, 2)
	fmt.Println(res)
}
