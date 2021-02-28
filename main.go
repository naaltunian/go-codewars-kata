package main

import (
	"fmt"

	"naaltunian/github.com/go-codewars-kata/kata"
)

// to run a function individually, use kata.functionName. All functions match their file name in the kata directory
func main() {
	indices := kata.TwoSum([]int{1, 2, 3}, 4)
	fmt.Println(indices)
}
