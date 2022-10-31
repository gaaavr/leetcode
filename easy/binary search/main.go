package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 5, 6, 8, 9, 10}, 5))
}
func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	firstIdx := 0
	lastIdx := len(arr) - 1
	for firstIdx <= lastIdx {
		idx := (lastIdx + firstIdx) / 2
		if arr[idx] == target {
			return idx
		}
		if target < arr[idx] {
			lastIdx = idx - 1
		} else {
			firstIdx = idx + 1
		}
	}
	return -1
}
