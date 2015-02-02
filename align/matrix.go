package align

type ScoringMatrix struct {
	data    []int
	Width   int
	Height  int
	maxCell Coord
}

type Coord struct {
	X, Y int
}

func CreateScoringMatrix(width, height int) ScoringMatrix {
	return ScoringMatrix{
		data:    make([]int, width*height),
		Width:   width,
		Height:  height,
		maxCell: Coord{-1, -1},
	}
}

func (m *ScoringMatrix) Get(x, y int) int {
	return m.data[m.i(x, y)]
}

func (m *ScoringMatrix) Set(x, y, value int) {
	m.data[m.i(x, y)] = value
}

func (m *ScoringMatrix) i(x, y int) int {
	x = x % m.Width
	y = y % m.Height
	return x + (m.Width * y)
}

func (m *ScoringMatrix) Max() (int, Coord) {
	if m.maxCell.X == -1 {
		max := 0
		for i := 0; i < m.Width; i++ {
			for j := 0; j < m.Height; j++ {
				if m.Get(i, j) >= max {
					m.maxCell = Coord{i, j}
					max = m.Get(i, j)
				}
			}
		}
	}
	return m.Get(m.maxCell.X, m.maxCell.Y), m.maxCell
}

func (m *ScoringMatrix) Traceback() []Coord {
	trace := make([]Coord, 1, m.Width)
	currentCell := m.maxCell
	trace[0] = currentCell
	for m.Get(currentCell.X, currentCell.Y) != 0 {
		// assume diagonal
		nextCell := Coord{currentCell.X - 1, currentCell.Y - 1}
		// check left
		if m.Get(nextCell.X, nextCell.Y) < m.Get(currentCell.X-1, currentCell.Y) {
			nextCell = Coord{currentCell.X - 1, currentCell.Y}
		}
		// check up
		if m.Get(nextCell.X, nextCell.Y) < m.Get(currentCell.X, currentCell.Y-1) {
			nextCell = Coord{currentCell.X, currentCell.Y - 1}
		}
		currentCell = nextCell
		trace = append(trace, nextCell)
	}
	return trace
}
