package leetcode

import (
	"fmt"
	"testing"
)

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5598 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}
func (s *Stack) Pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}
func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}
func (s *Stack) Len() int {
	return len(*s)
}
func trap(height []int) int {
	mStack := Stack{}
	res := 0
	for i, v := range height {
		if len(mStack) == 0 || height[mStack.Top()] >= v {
			mStack.Push(i)
			continue
		}
		prefixIndex := mStack.Pop()
		for len(mStack) > 0 && height[mStack.Top()] < v {
			res += (height[mStack.Top()] - height[prefixIndex]) * (i - mStack.Top() - 1)
			prefixIndex = mStack.Pop()
		}
		if mStack.Len() > 0 {
			res += (v - height[prefixIndex]) * (i - mStack.Top() - 1)
		}
		mStack.Push(i)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTrappingRainWater(t *testing.T) {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
