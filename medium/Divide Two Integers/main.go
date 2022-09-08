package main

import "fmt"

func main() {
	fmt.Println(divide(-2147483648, -1))
}

func divide(dividend int, divisor int) int {
	res := dividend / divisor
	if res < -1<<31 {
		return -1 << 31
	}
	if res > (1<<31)-1 {
		return (1 << 31) - 1
	}
	return res
}
