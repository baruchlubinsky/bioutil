package align

import (
	"fmt"
)

// http://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm#Example
func ExampleGlobalAlign() {
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
	alignmentA, alignmentB, matrix := GlobalAlign(seqA, seqB, gapPenalty, matchScore)
	fmt.Println(matrix.Max())
	fmt.Println(string(alignmentA))
	fmt.Println(string(alignmentB))
	// Output:
	// 12
	// A-CACACTA
	// AGCACAC-A
}
