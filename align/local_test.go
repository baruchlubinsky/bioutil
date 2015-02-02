package align

import (
	"fmt"
)

// http://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm#Example
func ExampleLocalAlign() {
	gapPenalty := func(x, y int) int {
		return -1
	}
	matchScore := func(a, b byte) int {
		if a == b {
			return 2
		} else {
			return -1
		}
	}
	seqA := []byte("ACACACTA")
	seqB := []byte("AGCACACA")
	alignmentA, alignmentB, _ := LocalAlign(seqA, seqB, gapPenalty, matchScore)
	fmt.Println(string(alignmentA))
	fmt.Println(string(alignmentB))
	// Output:
	// A-CACACTA
	// AGCACAC-A
}
