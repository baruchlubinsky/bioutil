package align

type ScoreFunc func(x, y int) int
type SimilarityFunc func(a, b byte) int

// Construct the scoring matrix for the (Smith-Waterman algorithm)[http://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm]
func BuildMatrix(seqA, seqB []byte, gapPenalty ScoreFunc, matchScore SimilarityFunc) ScoringMatrix {
	matrix := CreateScoringMatrix(len(seqA)+1, len(seqB)+1)
	matrix.Set(0, 0, 0)
	// Top row
	for i := 0; i < matrix.Width; i++ {
		matrix.Set(i+1, 0, 0)
	}
	// Left column
	for i := 0; i < matrix.Height; i++ {
		matrix.Set(0, i+1, 0)
	}
	for i := 1; i < matrix.Width; i++ {
		for j := 1; j < matrix.Height; j++ {
			score := 0
			max := 0
			// similarity
			if t := matrix.Get(i-1, j-1) + matchScore(seqA[i-1], seqB[j-1]); t > score {
				score = t
			}
			// deletion
			if t := matrix.Get(i-1, j) + gapPenalty(i, j); t > score {
				score = t
			}
			// insertion
			if t := matrix.Get(i, j-1) + gapPenalty(i, j); t > score {
				score = t
			}
			if score > max {
				matrix.maxCell = Coord{i, j}
				max = score
			}
			matrix.Set(i, j, score)
		}
	}
	return matrix
}

func ConstructAlignments(seqA, seqB []byte, traceback []Coord) (alignmentA, alignmentB []byte) {
	tempA := make([]byte, 0, len(seqA))
	tempB := make([]byte, 0, len(seqB))

	for i := 0; i < len(traceback)-1; i++ {
		cell := traceback[i]
		dx, dy := cell.X-traceback[i+1].X, cell.Y-traceback[i+1].Y
		if dx == 0 || cell.X > len(seqA) {
			tempA = append(tempA, '-')
		} else {
			tempA = append(tempA, seqA[cell.X-1])
		}
		if dy == 0 || cell.Y > len(seqB) {
			tempB = append(tempB, '-')
		} else {
			tempB = append(tempB, seqB[cell.Y-1])
		}
	}
	alignmentA = make([]byte, len(tempA))
	alignmentB = make([]byte, len(tempB))
	for i, b := range tempA {
		alignmentA[len(alignmentA)-1-i] = b
	}
	for i, b := range tempB {
		alignmentB[len(alignmentB)-1-i] = b
	}
	return
}

func LocalAlign(reference, query []byte, gapPenalty ScoreFunc, matchScore SimilarityFunc) (alignmentA, alignmentB []byte, maxScore int) {
	scores := BuildMatrix(reference, query, gapPenalty, matchScore)
	traceback := scores.Traceback()
	alignmentA, alignmentB = ConstructAlignments(reference, query, traceback)
	maxScore, _ = scores.Max()
	return
}
