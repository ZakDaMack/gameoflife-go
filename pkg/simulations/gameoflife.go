package simulations

import (
	"main/pkg/common"
	"math/rand"
)

type GameOfLife struct {
	seed   string
	queue  common.Queue[common.Canvas]
	canvas common.Canvas
}

func MakeGameOfLife(seed string) *GameOfLife {
	return &GameOfLife{
		seed:  seed,
		queue: common.MakeQueue[common.Canvas](3),
	}
}

func (g *GameOfLife) Start(height int, width int) common.Canvas {
	g.canvas = common.MakeCanvas(height, width)
	switch g.seed {
	case "glider":
		seedGlider(g.canvas)
	default:
		seedRandom(g.canvas)
	}
	g.queue.Add(g.canvas)
	return g.canvas
}

func (g *GameOfLife) Step() common.Canvas {
	newWorld := common.MakeCanvas(g.canvas.Size())
	for i := range g.canvas {
		for j := range g.canvas[i] {
			newWorld[i][j] = nextCellState(j, i, g.canvas)
		}
	}
	g.canvas = newWorld
	g.queue.Add(newWorld)
	return newWorld
}

func (g *GameOfLife) ShouldStop() bool {
	q := &g.queue
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
