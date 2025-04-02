package leetcode

import (
	"fmt"
	"math"
	"testing"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//

// leetcode submit region begin(Prohibit modification and deletion)
type stack []int

func (s *stack) Top() int {
	return (*s)[len(*s)-1]
}
func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Pop() {
	*s = (*s)[:len(*s)-1]
}
func largestRectangleArea(heights []int) int {
	heights = append(heights, math.MinInt)
	mStack := stack{}
	res := 0
	for i, v := range heights {
		if len(mStack) == 0 || heights[mStack.Top()] < v {
			mStack.Push(i)
		} else {
			for len(mStack) > 0 && heights[mStack.Top()] > v {
				height := heights[mStack.Top()]
				mStack.Pop()
				leftIndex := -1
				if len(mStack) > 0 {
					leftIndex = mStack.Top()
				}
				area := (i - leftIndex - 1) * height
				res = max(area, res)
			}
			mStack.Push(i)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLargestRectangleInHistogram(t *testing.T) {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(nums))
}
