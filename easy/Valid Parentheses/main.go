package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Every close bracket has a corresponding open bracket of the same type.

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
