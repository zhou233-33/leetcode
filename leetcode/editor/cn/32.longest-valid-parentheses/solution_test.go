package leetcode

import (
	"fmt"
	"testing"
)

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
// Related Topics 栈 字符串 动态规划 👍 2661 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	// longest valida string ends with i chart
	if len(s) <= 1 {
		return 0
	}
	f := make([]int, len(s))
	f[0] = 0
	res := 0
	var leftChart byte = '('
	for i := 1; i < len(s); i++ {
		if s[i] == leftChart {
			f[i] = 0
			continue
		}
		//match i-1
		match1 := 0
		match2 := 0
		if s[i-1] == leftChart {
			if i-2 >= 0 {
				match1 += f[i-2]
			}
			match1 += 1
		}
		if left := i - 1 - f[i-1]*2; left >= 0 {
			if s[left] == leftChart {
				match2 = f[i-1] + 1
				if left-1 >= 0 {
					match2 += f[left-1]
				}
			}
		}
		f[i] = max(match1, match2)
		res = max(f[i], res)
	}
	return res * 2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses(""))
}
