package leetcode

import (
	"fmt"
	"testing"
)

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 3791 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	type str struct {
		Val        string //记录值
		CountLeft  int    //记录左括号的次数
		CountRight int    //记录右括号的次数
	}
	l1 := make([]str, 0)
	l1 = append(l1, str{
		Val:        "",
		CountLeft:  0,
		CountRight: 0,
	})
	for i := 0; i < 2*n; i++ {
		for len(l1[0].Val) == i {
			valStr := l1[0].Val
			if l1[0].CountLeft < n {
				l1 = append(l1, str{
					Val:        valStr + "(",
					CountLeft:  l1[0].CountLeft + 1,
					CountRight: l1[0].CountRight,
				})
			}
			if l1[0].CountRight < n && l1[0].CountLeft > l1[0].CountRight {
				l1 = append(l1, str{
					Val:        valStr + ")",
					CountLeft:  l1[0].CountLeft,
					CountRight: l1[0].CountRight + 1,
				})
			}
			l1 = l1[1:]
		}
	}
	res := make([]string, 0)
	for _, item := range l1 {
		res = append(res, item.Val)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerateParentheses(t *testing.T) {
	for _, item := range generateParenthesis(3) {
		fmt.Println(item)
	}
}
