package kata

import "testing"

func TestComplimentaryDNA(t *testing.T) {
	x := DNAStrand("ATTGC")
	if x != "TAACG" {
		t.Errorf("Expected TAACG but got %s", x)
	}

	y := DNAStrand("GTAT")
	if y != "CATA" {
		t.Errorf("Expected CATA but got %s", y)
	}
}
