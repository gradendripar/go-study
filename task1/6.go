package main

/**
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。元素的相对顺序应该保持一致 。然后返回 nums 中唯一元素的个数。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length] {
			length++
			nums[length] = nums[i]
		}
	}
	return length + 1
}
