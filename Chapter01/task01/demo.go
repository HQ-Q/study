package main

import (
	"sort"
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

// 最长公共前缀 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 加一 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
// 思路
// 1. 从数组末尾开始遍历
// 2. 如果当前位不是9，直接加一并返回
// 3. 如果是9，将其置为0，继续处理前一位
// 4. 如果所有位都是9，需要扩展数组长度，在开头添加1
// 例如：[1,2,3] → [1,2,4]，[9,9] → [1,0,0]
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// 删除有序数组中的重复项
// 思路：使用两个指针，一个指向当前位置，一个指向下一个位置
// 1. 如果当前位置的值等于下一个位置的值，则跳过下一个位置
// 2. 如果当前位置的值不等于下一个位置的值，则将当前位置的值赋给下一个位置
// 3. 返回当前位置的下一个位置
func removeDuplicates(nums []int) (int, []int) {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	n := i + 1
	//删除数组nums[n]后面的元素
	nums = nums[:n]
	return n, nums
}

// 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，
// 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
// 如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

// 两数之和 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) (int, int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return nums[i], nums[j]
			}
		}
	}
	return 0, 0
}

func main() {
	//nums := []int{1, 2, 4, 1, 2}
	//fmt.Println(singleNumber(nums))

	//fmt.Println(isValid("()]"))

	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	//fmt.Println(plusOne([]int{1, 2, 3}))
	//fmt.Println(plusOne([]int{4, 3, 2, 1}))

	//fmt.Println(removeDuplicates([]int{1, 1, 2}))
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
