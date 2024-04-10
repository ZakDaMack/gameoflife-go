package simulations

import (
	"main/pkg/common"
	"math/rand"
)

func Start(height, width int) common.Canvas {
	c := common.MakeCanvas(height, width)
	seedRandom(c)
	// seedGlider(c)
	return c
}

func Step(c common.Canvas) common.Canvas {
	newWorld := common.MakeCanvas(c.Size())
	for i := range c {
		for j := range c[i] {
			newWorld[i][j] = nextCellState(j, i, c)
		}
	}
	return newWorld
}

func ShouldStop(q common.Queue[common.Canvas]) bool {
	for i := 0; i < q.Length()-1; i++ {
		for j := i + 1; j < q.Length(); j++ {
			if q.Get(i).IsSame(q.Get(j)) {
				return true
			}
		}
	}

	return false
}

func activeNeighbours(x, y int, c common.Canvas) int {
	n := 0
	// go through neighbours and get status
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if c.GetCell(j, i) {
				n++
			}
		}
	}

	if c.GetCell(x, y) {
		n--
	}

	return n
}

func nextCellState(x, y int, c common.Canvas) bool {
	n := activeNeighbours(x, y, c)
	s := c.GetCell(x, y)
	return (n == 3) || (n == 2 && s)
}

func seedRandom(c common.Canvas) {
	for _, row := range c {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func seedGlider(c common.Canvas) {
	d := [][]bool{
		{},
		{false, false, true, false, true},
		{false, false, false, true, true},
		{false, false, false, true, false},
	}

	for y, row := range d {
		for x, val := range row {
			c[y][x] = val
		}
	}
}
