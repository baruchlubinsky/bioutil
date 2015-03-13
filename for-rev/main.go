package main

import (
	"fmt"
	"hyraxbio.kilnhg.com/golang/bioutil.git"
	"hyraxbio.kilnhg.com/golang/bioutil.git/align"
	"hyraxbio.kilnhg.com/golang/seq2res.git/data"
	"os"
	"strings"
)

func main() {
	query := []byte(os.Args[1])
	reference := []byte(strings.ToUpper(data.Hxb2))

	gapPenalty := func(x, y int) int {
		if y == 0 {
			return 0
		}
		return -1
	}
	matchScore := func(a, b byte) int {
		if a == b {
			return 2
		} else {
			return -1
		}
	}

	reverse := bioutil.ReverseCompliment(query)

	_, _, forwardScores := align.GlobalAlign(reference, query, gapPenalty, matchScore)
	_, _, reverseScores := align.GlobalAlign(reference, reverse, gapPenalty, matchScore)

	fwd, _ := forwardScores.Max()
	rev, _ := reverseScores.Max()

	if fwd > rev {
		fmt.Println("Forward")
		fmt.Printf("%v > %v\n", fwd, rev)
	} else {
		fmt.Println("Reverse")
		fmt.Printf("%v > %v\n", rev, fwd)
	}
}
