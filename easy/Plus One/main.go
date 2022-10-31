package main

import (
	"fmt"
)

// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least
// significant in left-to-right order. The large integer does not contain any leading 0's.
//
//Increment the large integer by one and return the resulting array of digits.

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	var ind bool
	if digits[len(digits)-1]+1 == 10 {
		for i := len(digits) - 1; i >= 0; i-- {
			if i == len(digits)-1 {
				if len(digits) == 1 {
					return []int{1, 0}
				}
				digits[i] = 0
				ind = true
				continue
			}
			if ind {
				if digits[i] == 9 && i != 0 {
					digits[i] = 0
					continue
				}
				digits[i]++
				if digits[i] != 0 {
					ind = false
				}
			}
		}
	} else {
		digits[len(digits)-1]++
		return digits
	}
	if digits[0] == 10 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
