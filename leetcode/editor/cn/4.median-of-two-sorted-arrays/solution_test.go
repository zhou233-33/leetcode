package leetcode

import (
	"fmt"
	"testing"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics 数组 二分查找 分治 👍 7417 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//左边的 index
	left1 := 0
	left2 := 0
	right1 := len(nums1) - 1
	right2 := len(nums2) - 1
	//是否遍历完数组的标志
	flag1 := false
	flag2 := false
	var left1Value, right1Value, left2Value, right2Value int
	var midVal float64
	for {
		var minVal, maxVal int
		if left1 > right1 {
			flag1 = true
		} else {
			left1Value = nums1[left1]
			right1Value = nums1[right1]
		}
		if left2 > right2 {
			flag2 = true
		} else {
			left2Value = nums2[left2]
			right2Value = nums2[right2]
		}
		if flag1 && flag2 {
			return midVal
		}
		if flag2 {
			minVal = left1Value
			maxVal = right1Value
			left1++
			right1--
		}
		if flag1 {
			minVal = left2Value
			maxVal = right2Value
			left2++
			right2--
		}
		if !flag1 && !flag2 {
			if left1Value < left2Value {
				minVal = left1Value
				left1++
			} else {
				minVal = left2Value
				left2++
			}
			if right1Value > right2Value {
				maxVal = right1Value
				right1--
			} else {
				maxVal = right2Value
				right2--
			}
		}
		midVal = (float64(minVal) + float64(maxVal)) / 2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	l1 := []int{1, 2}
	l2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(l1, l2))
}
