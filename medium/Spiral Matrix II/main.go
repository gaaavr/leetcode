package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for arr := range matrix {
		matrix[arr] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	var count = 1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			matrix[top][i] = count
			count++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = count
			count++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}
	return matrix
}
