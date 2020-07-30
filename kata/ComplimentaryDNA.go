package kata

// INSTRUCTIONS: In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". You have function with one side of the DNA (string, except for Haskell); you need to get the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

func DNAStrand(dna string) string {
	var dnaCompliments = map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}
	var compliments string

	for _, char := range dna {
		compliments += dnaCompliments[string(char)]
	}

	return compliments
}
