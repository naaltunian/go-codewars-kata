package kata

// INSTRUCTIONS: given a postive int n, return the number of divisors

func Divisors(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}
