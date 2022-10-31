package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order,
//remove the duplicates in-place such that each unique element appears only once.
//The relative order of the elements should be kept the same.

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	k := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			k--
			i--
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return k
}
