package main

import (
	"strings"
)

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	var count, prevCount string
	if len(strs) == 1 {
		return strs[0]
	}
	var min int
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" || strs[i-1] == "" {
			return ""
		}
		if len(strs[i]) < len(strs[i-1]) {
			min = len(strs[i])
		} else {
			min = len(strs[i-1])
		}
		for j := 0; j < min; j++ {
			if i == 1 {
				if strs[i][j] != strs[i-1][j] {
					prevCount = count
					break
				}
				if j == len(strs[i])-1 {
					count += string(strs[i][j])
					prevCount = count
					break
				}
				count += string(strs[i][j])
				continue
			}
			if strings.HasPrefix(prevCount, count+string(strs[i][j])) {
				count += string(strs[i][j])
				continue
			} else {
				break
			}
		}
		prevCount = count
		count = ""
	}
	return prevCount
}
