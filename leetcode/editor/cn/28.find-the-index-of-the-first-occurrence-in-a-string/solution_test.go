package leetcode

import (
	"fmt"
	"testing"
)

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
//如果 needle 不是 haystack 的一部分，则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
//
//
// 示例 2：
//
//
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack 和 needle 仅由小写英文字符组成
//
//
// Related Topics 双指针 字符串 字符串匹配 👍 2375 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	transform := calculate(needle)
	i := 0
	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[i] {
			i++
			if i == len(needle) {
				return j - len(needle) + 1
			}
		} else {
			//jump
			for i != -1 && needle[i] != haystack[j] {
				i = transform[i]
			}
			// if i==-1 means the first element is not equal too,
			// we need to check the next item of haystack whether equal to the first element
			// if needle[i] == haystack[j] , we don't need to move i
			i++
		}
	}
	return -1
}
func calculate(needle string) []int {
	transform := make([]int, len(needle))
	left := 0
	transform[0] = -1
	cnt := 0
	for i := 1; i < len(needle); i++ {
		transform[i] = cnt
		if needle[left] == needle[i] {
			left++
			cnt++
		} else {
			for left > -1 && needle[left] != needle[i] {
				left = transform[left]
			}
			cnt = left + 1
			left++
		}
	}
	return transform
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	haystack := "aabaaabaaac"
	needle := "aabaaac"
	fmt.Println(strStr(haystack, needle))
}
