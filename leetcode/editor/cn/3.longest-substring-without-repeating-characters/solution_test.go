package leetcode

import (
	"fmt"
	"testing"
)

//Given a string s, find the length of the longest substring without duplicate
//characters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10642 👎 0

/*
解题思路
从左向右进行遍历，维护一个 map 存储已经遍历过的字符以及对应的下标。
维护一个变量表示字符串的起点，通过获取字符的下标是否大于起点来判断字符是否重复。
还需要维护一个变量来记录最大字符串的长度
*/
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	valIndex := make(map[byte]int)
	maxLen := 0
	left := 0
	for i := 0; i < len(s); i++ {
		letter := s[i]
		if index, ok := valIndex[letter]; ok && index >= left {
			currentLen := i - left
			if currentLen > maxLen {
				maxLen = currentLen
			}
			left = index + 1
		}
		valIndex[letter] = i
	}
	currentLen := len(s) - left
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
