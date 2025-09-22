package main

import "fmt"

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func singleNumber(nums []int) int {
	var res = 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	nums := []int{1, 2, 4, 1, 2}
	fmt.Println(singleNumber(nums))
}
