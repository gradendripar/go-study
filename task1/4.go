package main

/**
14. 最长公共前缀：编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
