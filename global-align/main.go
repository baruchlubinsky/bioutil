package main

import (
	"fmt"
	"code.hyraxbio.co.za/bioutil/align"

	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(os.Args)
		log.Fatalln("Align two sequences. Input is two strings provided as arguments to this program.")
	}

	querySequence := []byte(os.Args[1])
	referenceSequence := []byte(os.Args[2])

	fmt.Println(MidCompare(string(referenceSequence), string(querySequence)))
}

func MidCompare(reference, query string) int {
	gapPenalty := func(x, y int) int {
		// if y > len(query) {
		// 	return 0
		// }
		return -1
	}
	matchScore := func(a, b byte) int {
		if a == b {
			return 2
		} else {
			return -1
		}
	}
	// Implement a traceback function matching Ram's
	traceback := func(m *align.ScoringMatrix) []align.Coord {
		trace := make([]align.Coord, 1, m.Width)
		currentCell := align.Coord{m.Width - 1, m.Height - 1}
		trace[0] = currentCell
		for currentCell.X > 0 && currentCell.Y > 0 {
			pos := m.Get(currentCell.X, currentCell.Y)
			var nextCell align.Coord
			// diagonal
			if m.Get(currentCell.X-1, currentCell.Y-1)+matchScore(reference[currentCell.X-1], query[currentCell.Y-1]) == pos {
				nextCell = align.Coord{currentCell.X - 1, currentCell.Y - 1}
			} else if m.Get(currentCell.X-1, currentCell.Y)+gapPenalty(0, 0) == pos { //left
				nextCell = align.Coord{currentCell.X - 1, currentCell.Y}
			} else if m.Get(currentCell.X, currentCell.Y-1)+gapPenalty(0, 0) == pos { // up
				nextCell = align.Coord{currentCell.X, currentCell.Y - 1}
			} else {
				break
			}
			fmt.Println(currentCell.String() + " -> " + nextCell.String())

			currentCell = nextCell

			trace = append(trace, nextCell)
		}
		for currentCell.X > 0 {
			trace = append(trace, align.Coord{currentCell.X - 1, currentCell.Y})
			currentCell.X--
		}
		for currentCell.Y > 0 {
			trace = append(trace, align.Coord{currentCell.X, currentCell.Y - 1})
			currentCell.Y--
		}
		return trace
	}

	// alignment1, alignment2, scores := align.GlobalAlign([]byte(reference), []byte(query), gapPenalty, matchScore)
	scores := align.BuildMatrix([]byte(reference), []byte(query), gapPenalty, matchScore)
	trace := traceback(&scores)
	alignment1, alignment2 := align.ConstructAlignments([]byte(reference), []byte(query), trace)

	fmt.Println(trace)

	fmt.Println(scores.String())

	fmt.Println(string(alignment1))
	fmt.Println(string(alignment2))

	score, _ := scores.Max()
	if score == 2*len(query) {
		return 0
	} else {
		res := len(alignment2)
		for i := range alignment2 {
			if alignment1[i] == alignment2[i] {
				res--
			}
			// Because gaps are included in the original value.
			// if alignment2[i] == '-' {
			// 	res--
			// }
		}
		return res
	}
}
