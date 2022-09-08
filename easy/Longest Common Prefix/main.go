package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("[[]]"))
}

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	arr := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			arr = append(arr, v)
			continue
		} else {
			lastIdx := len(arr) - 1
			if lastIdx == -1 {
				return false
			}
			if arr[lastIdx] == m[v] {
				arr = arr[:lastIdx]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}
