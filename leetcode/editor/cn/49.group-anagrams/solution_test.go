package leetcode

import (
	"fmt"
	"testing"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
//
// Related Topics 数组 哈希表 字符串 排序 👍 2197 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, item := range strs {
		str := generateStr(item)
		strMap[str] = append(strMap[str], item)
	}
	res := make([][]string, 0)
	for _, val := range strMap {
		res = append(res, val)
	}
	return res
}

func generateStr(s string) string {
	chartCnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chart := s[i]
		chartCnt[chart-'a']++
	}
	var res string
	for index, cnt := range chartCnt {
		if cnt == 0 {
			continue
		}
		res += fmt.Sprintf("%d%c", cnt, index+'a')
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGroupAnagrams(t *testing.T) {
	str := []string{""}
	fmt.Println(groupAnagrams(str))
}
