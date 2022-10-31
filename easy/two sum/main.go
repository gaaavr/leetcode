package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2, 2)
	res := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := res[target-v]; ok {
			ans[0] = res[target-v]
			ans[1] = i
			return ans
		}
		res[v] = i
	}
	return nil
}
