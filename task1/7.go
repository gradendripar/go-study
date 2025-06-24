package main

/**
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
链接：https://leetcode-cn.com/problems/merge-intervals/
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			merged = append(merged, curr)
		}
	}
	return merged
}
