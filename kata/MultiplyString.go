package kata

import (
	"fmt"
	"strconv"
)

// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
func MultiplyString(num1 string, num2 string) string {
	int1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println(err)
		return "0"
	}

	int2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err)
		return "0"
	}

	if int1 == 0 || int2 == 0 {
		return "0"
	}

	val := int1 * int2
	stringVal := strconv.Itoa(val)

	return stringVal
}
