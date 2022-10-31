package main

import (
	"fmt"
	"math"
)

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// The returned integer should be non-negative as well.

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
