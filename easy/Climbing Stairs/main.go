package main

import "fmt"

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	prevNumber := 1
	currentNumber := 1
	for i := 1; i <= n; i++ {
		if i < 2 {
			continue
		} else {
			currentNumber += prevNumber
			prevNumber = currentNumber - prevNumber
		}
	}
	return currentNumber
}
