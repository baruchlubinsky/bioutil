package align

import "fmt"

type ScoringMatrix struct {
	data    []int
	Width   int
	Height  int
	maxCell *Coord
}

type Coord struct {
	X, Y int
}

// a - b
func (a *Coord) Diff(b Coord) Coord {
	return Coord{a.X - b.X, a.Y - b.Y}
}

// a == b
func (a *Coord) Equal(b Coord) bool {
	return a.X == b.X && a.Y == b.Y
}

func (c *Coord) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}

func (m *ScoringMatrix) String() string {
	res := ""
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			res += fmt.Sprintf("%v\t", m.Get(j, i))
		}
		res += "\n"
	}
	return res
}

func CreateScoringMatrix(width, height int) ScoringMatrix {
	return ScoringMatrix{
		data:   make([]int, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *ScoringMatrix) Get(x, y int) int {
	if x >= m.Width || y >= m.Height || m.i(x, y) >= len(m.data) {
		fmt.Printf("accessing (%v,%v) in [%v,%v]\n", x, y, m.Width, m.Height)
	}
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
	if m.maxCell == nil {
		max := 0
		for i := 0; i < m.Width; i++ {
			for j := 0; j < m.Height; j++ {
				if m.Get(i, j) >= max {
					m.maxCell = &Coord{i, j}
					max = m.Get(i, j)
				}
			}
		}
	}
	return m.Get(m.maxCell.X, m.maxCell.Y), *m.maxCell
}

// Starts from the highest scoring cell in the ScoringMatrix and proceeds to an edge.
func (m *ScoringMatrix) Traceback() []Coord {
	trace := make([]Coord, 1, m.Width)
	_, currentCell := m.Max()
	trace[0] = currentCell
	for currentCell.X > 0 && currentCell.Y > 0 {
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
