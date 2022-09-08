package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{}))
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
