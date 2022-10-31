package main

import "fmt"

// Given a string s consisting of words and spaces, return the length of the last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.
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
