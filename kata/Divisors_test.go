package kata

import "testing"

func TestDivisors(t *testing.T) {
	if 3 != Divisors(4) {
		t.Errorf("Divisors(4) returned incorrect value")
	}
}
