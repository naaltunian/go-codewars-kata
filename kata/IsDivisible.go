package kata

// INSTRUCTIONS: Create a function that checks if a number n is divisible by two numbers x AND y. All inputs are positive, non-zero digits.

func IsDivisible(n, x, y int) bool {
	// multi-line
	// if n%x == 0 && n%y == 0 {
	// 	return true
	// }

	// return false

	// single-line
	return n%x == 0 && n%y == 0
}
