package align

import (
	"testing"
)

func TestCreate(t *testing.T) {
	w := 5
	h := 7
	m := CreateScoringMatrix(w, h)
	if m.Width != w {
		t.Error("Incorrect width")
	}
	if m.Height != h {
		t.Error("Incorrect height")
	}
}

func TestAssignment(t *testing.T) {
	w := 5
	h := 7
	m := CreateScoringMatrix(w, h)
	f := func(x, y int) int {
		return x*y + y
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			m.Set(i, j, f(i, j))
		}
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if m.Get(i, j) != f(i, j) {
				t.Errorf("Wrong value at (%v, %v), got %v instead of %v", i, j, m.Get(i, j), f(i, j))
			}
		}
	}
}
