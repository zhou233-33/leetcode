package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
// 输入：s = "()"
//
//
// 输出：true
//
// 示例 2：
//
//
// 输入：s = "()[]{}"
//
//
// 输出：true
//
// 示例 3：
//
//
// 输入：s = "(]"
//
//
// 输出：false
//
// 示例 4：
//
//
// 输入：s = "([])"
//
//
// 输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 4664 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	deleteMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	l := list.New()
	for i := 0; i < len(s); i++ {
		item := s[i]
		if l.Len() == 0 {
			l.PushBack(item)
			continue
		}
		corItem, ok := deleteMap[item]
		if !ok {
			l.PushBack(item)
			continue
		}
		if l.Back().Value == corItem {
			l.Remove(l.Back())
			continue
		}
		l.PushBack(item)
	}
	return l.Len() == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {
	s := "{([]()){}}"
	fmt.Println(isValid(s))
}
