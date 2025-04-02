package leetcode

import (
	"fmt"
	"testing"
)

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//
//
// 示例 1：
//
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
//
// 示例 2：
//
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//     注意，你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅由小写英文字母组成
// wordDict 中的所有字符串 互不相同
//
//
// Related Topics 字典树 记忆化搜索 数组 哈希表 字符串 动态规划 👍 2718 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	wordMap := make(map[string]bool)
	for _, item := range wordDict {
		wordMap[item] = true
	}
	for i := 0; i < len(dp); i++ {
		subStr := s[:i+1]
		if wordMap[subStr] == true {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if dp[j] == true {
				backStr := s[j+1 : i+1]
				if wordMap[backStr] == true {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(dp)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordBreak(t *testing.T) {
	strsDict := []string{
		"leet", "cod", "e"}
	s := "leetcode"
	fmt.Println(wordBreak(s, strsDict))
}
