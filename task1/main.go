package main

import (
	"fmt"
)

func main() {
	// 1. 只出现一次的数字
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))

	// 2. 回文数
	num := 121
	fmt.Println(isPalindrome(num))

	// 3. 有效的括号
	s := "()"
	fmt.Println(isValid(s))

	// 4. 最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	// 5. 加一
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))

	// 6. 删除有序数组中的重复项
	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

	// 7. 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	// 8. 两数之和
	nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
