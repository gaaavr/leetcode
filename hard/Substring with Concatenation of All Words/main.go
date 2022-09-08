package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("ababababab", []string{"ababa", "babab"}))
}

func findSubstring(s string, words []string) []int {
	arr := make([]int, 0, 0)
	l := len(words) * len(words[0])
	for i := 0; i < len(s); i++ {
		var idx bool
		if i+l > len(s) {
			return arr
		}
		subs := s[i : i+l]
		store := make(map[string]int)
		for _, v := range words {
			store[v]++
		}
		for j := 0; j < len(subs); j += len(words[0]) {
			if _, ok := store[subs[j:j+len(words[0])]]; !ok {
				idx = true
				break
			}
			store[subs[j:j+len(words[0])]]--
		}
		if !idx {
			for _, v := range store {
				if v != 0 {
					idx = true
					break
				}
			}
		}
		if !idx {
			arr = append(arr, i)
		}
	}
	return arr
}
