package kata

import (
	"fmt"
	"strconv"
)

func Reverse(num int) int {
	i := strconv.Itoa(num)

	fmt.Println(i)
	fmt.Println(i[0])

	runes := []rune(i)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str := string(runes)

	ints, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	return ints
}
