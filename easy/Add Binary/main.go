package main

import (
	"fmt"
	"math/big"
)

// Given two binary strings a and b, return their sum as a binary string.

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	numA := new(big.Int)
	numB := new(big.Int)
	numA.SetString(a, 2)
	numB.SetString(b, 2)
	return new(big.Int).Add(numA, numB).Text(2)
}
