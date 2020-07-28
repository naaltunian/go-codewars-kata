package kata

import "testing"

func TestIsDivisible(t *testing.T) {
	if !IsDivisible(3, 1, 3) {
		t.Error("Failed")
	}

	if !IsDivisible(12, 2, 6) {
		t.Error("Failed")
	}

	if IsDivisible(100, 5, 3) {
		t.Error("Failed")
	}

	if IsDivisible(12, 7, 5) {
		t.Error("Failed")
	}
}
