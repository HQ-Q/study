package main

import (
	"fmt"
	"strings"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	length := len(s) / 2
	for i := 0; i < length; i++ {
		s = strings.Replace(s, "()", "", 1)
		s = strings.Replace(s, "[]", "", 1)
		s = strings.Replace(s, "{}", "", 1)
	}
	return len(s) == 0
}

func main() {
	s := "([{}])"
	fmt.Println(isValid(s))
	print("Hello World")
}
