package gameoflife

import (
	"fmt"
	"math/rand"
)

type World [][]bool

func MakeWorld(height, width int) World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}

	return w
}

func (w World) Seed() {
	for _, row := range w {
		for i := range row {
			if rand.Intn(3) == 1 {
				row[i] = true
			}
		}
	}
}

func (w World) Print() {
	for _, row := range w {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf("\xF0\x9F\x9F\xA9")
			default:
				fmt.Printf("\xF0\x9F\x9F\xAB")
			}
		}
		fmt.Printf("\n")
	}
}

func (w World) size() (int, int) {
	// returns height, width
	return len(w), len(w[0])
}

func (w World) Copy() World {
	newWorld := MakeWorld(w.size())
	for i := range w {
		copy(newWorld[i], w[i])
	}
	return newWorld
}

func (a World) IsSame(b World) bool {
	for i := range a {
		for j := range a[i] {
			if b[i][j] != a[i][j] {
				return false
			}
		}
	}
	return true
}

func (w World) Step() World {
	newWorld := MakeWorld(w.size())
	for i := range w {
		for j := range w[i] {
			newWorld[i][j] = w.nextCellState(i, j)
		}
	}
	return newWorld
}

func (w World) status(x, y int) bool {
	height, width := w.size()
	b := (height + y) % height
	a := (width + x) % width
	return w[b][a]
}

func (w World) activeNeighbours(x, y int) int {
	n := 0
	// go through neighbours and get status
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if w.status(i, j) {
				n++
			}
		}
	}

	if w.status(x, y) {
		n--
	}

	return n
}

func (w World) nextCellState(x, y int) bool {
	// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
	// Any live cell with two or three live neighbors lives on to the next generation.
	// Any live cell with more than three live neighbors dies, as if by overpopulation.
	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	n := w.activeNeighbours(x, y)
	s := w.status(x, y)
	return (n == 3) || (n == 2 && s)
}
