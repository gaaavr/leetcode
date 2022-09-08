package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 9}, 8))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var i int
	first := 0
	last := len(nums) - 1
	for first <= last {
		i = (first + last) / 2
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			last = i - 1
		} else {
			first = i + 1
		}
	}
	return first
}
