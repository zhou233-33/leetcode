package leetcode

import (
	"fmt"
	"testing"
)

//给你一个字符串 s，找到 s 中最长的 回文 子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 双指针 字符串 动态规划 👍 7552 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var res string
	var maxLen int
	travel := func(left, right int) (int, string) {
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		return right - left - 1, s[left+1 : right]
	}
	for i := 0; i < len(s); i++ {
		//偶数中心
		left := i
		right := i + 1
		strLen, str := travel(left, right)
		if strLen > maxLen {
			maxLen = strLen
			res = str
		}
		//奇数中心
		left = i - 1
		strLen, str = travel(left, right)
		if strLen > maxLen {
			maxLen = strLen
			res = str
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestPalindromicSubstring(t *testing.T) {
	s := "abcdbdcb"
	fmt.Println(longestPalindrome(s))
}
