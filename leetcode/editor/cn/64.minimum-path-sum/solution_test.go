package leetcode

import (
	"fmt"
	"math"
	"testing"
)

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 200
//
//
// Related Topics 数组 动态规划 矩阵 👍 1792 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
				continue
			}
			minStep := math.MaxInt
			if i > 0 {
				minStep = dp[i-1][j]
			}
			if j > 0 {
				minStep = min(dp[i][j-1], minStep)
			}
			dp[i][j] = minStep + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumPathSum(t *testing.T) {
	grid := [][]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(minPathSum(grid))
}
