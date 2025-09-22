package main

import (
	"strings"
)

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func singleNumber(nums []int) int {
	var res = 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

// 有效的括号 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	s = strings.ReplaceAll(s, "{}", "")
	s = strings.ReplaceAll(s, "()", "")
	s = strings.ReplaceAll(s, "[]", "")
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	//nums := []int{1, 2, 4, 1, 2}
	//fmt.Println(singleNumber(nums))

	//fmt.Println(isValid("()]"))
}
