package main

import (
	"fmt"

	"naaltunian/github.com/go-codewars-kata/kata"
)

// to run a function individually, use kata.functionName. All functions match their file name in the kata directory
func main() {
	x := kata.IsDivisible(7, 12, 5)
	fmt.Println(x)
}
