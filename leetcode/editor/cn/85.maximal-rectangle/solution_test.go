package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
// ,["1","0","0","1","0"]]
// 输出：6
// 解释：最大矩形如上图所示。
//
// 示例 2：
//
// 输入：matrix = [["0"]]
// 输出：0
//
// 示例 3：
//
// 输入：matrix = [["1"]]
// 输出：1
//
// 提示：
//
// rows == matrix.length
// cols == matrix[0].length
// 1 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'
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
func maximalRectangle(matrix [][]byte) int {
	CountElements(matrix)
	res := 0
	calculateArea := func(mStack *stack, rightIndex, rightVal, j int) {
		for len(*mStack) > 0 && ConvertByteToInt(matrix[mStack.Top()][j]) > rightVal {
			height := ConvertByteToInt(matrix[mStack.Top()][j])
			mStack.Pop()
			leftIndex := -1
			if len(*mStack) >= 1 {
				leftIndex = mStack.Top()
			}
			area := (rightIndex - leftIndex - 1) * height
			res = max(area, res)
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		mStack := stack{}
		for i := 0; i < len(matrix); i++ {
			if len(mStack) == 0 || ConvertByteToInt(matrix[mStack.Top()][j]) < ConvertByteToInt(matrix[i][j]) {
				mStack.Push(i)
				continue
			} else {
				rightIndex := i
				rightVal := ConvertByteToInt(matrix[i][j])
				calculateArea(&mStack, rightIndex, rightVal, j)
				mStack.Push(i)
			}
		}
		if len(mStack) != 0 {
			rightIndex := len(matrix)
			rightVal := math.MinInt32
			calculateArea(&mStack, rightIndex, rightVal, j)
		}
	}
	return res
}
func ConvertByteToInt(i byte) int {
	return int(i - '0')
}
func CountElements(m [][]byte) {
	for _, item := range m {
		cnt := 0
		for i := len(item) - 1; i >= 0; i-- {
			if item[i] == '1' {
				cnt++
				item[i] = byte('0' + cnt)
			} else {
				cnt = 0
			}
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
func TestMaximalRectangle(t *testing.T) {
	test := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	res := maximalRectangle(test)
	fmt.Println(res)
}
