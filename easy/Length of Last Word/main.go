package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord(""))
}

func lengthOfLastWord(s string) int {
	var ind bool
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if ind {
				return count
			}
			continue
		}
		ind = true
		count++
	}
	return count
}
