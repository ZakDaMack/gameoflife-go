package main

import (
	"fmt"
	"math/rand"
	"time"
	// "example.com/gameoflife"
)

const (
	MAX_LENGTH = 3
)

type World [][]bool
type WorldHistory []World

func (h WorldHistory) Add(w World) WorldHistory {
	// append copy to the queue
	h = append(h, w)
	// if the queue is over length, remove the first item
	if len(h) > MAX_LENGTH {
		h = h[1:]
	}
	return h
}

func MakeWorld(height int, width int) World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}

	return w
}

func (w World) Seed() {
	for _, row := range w {
		for i := range row {
			if rand.Intn(4) == 1 {
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

func (w World) Size() (int, int) {
	// returns height, width
	return len(w), len(w[0])
}

func (w World) Status(x, y int) bool {
	height, width := w.Size()
	b := (height + y) % height
	a := (width + x) % width
	return w[b][a]
}

func (w World) Neighbours(x, y int) int {
	n := 0
	// go through neighbours and get status
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if w.Status(i, j) {
				n++
			}
		}
	}

	if w.Status(x, y) {
		n--
	}

	return n
}

func (w World) NextState(x, y int) bool {
	// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
	// Any live cell with two or three live neighbors lives on to the next generation.
	// Any live cell with more than three live neighbors dies, as if by overpopulation.
	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	n := w.Neighbours(x, y)
	s := w.Status(x, y)
	return (n == 3) || (n == 2 && s)
}

func (w World) Step() World {
	newWorld := MakeWorld(w.Size())
	for i := range w {
		for j := range w[i] {
			newWorld[i][j] = w.NextState(i, j)
		}
	}
	return newWorld
}

func (w World) Copy() World {
	newWorld := MakeWorld(w.Size())
	for i := range w {
		copy(newWorld[i], w[i])
	}
	return newWorld
}

func IsSame(a, b World) bool {
	for i := range a {
		for j := range a[i] {
			if b[i][j] != a[i][j] {
				return false
			}
		}
	}
	return true
}

func ContainsDuplicates(h WorldHistory) bool {
	for i := 0; i < len(h)-1; i++ {
		for j := i + 1; j < len(h); j++ {
			if IsSame(h[i], h[j]) {
				return true
			}
		}
	}

	return false
}

func main() {
	// i := 0
	world := MakeWorld(20, 20)
	h := make(WorldHistory, 0)
	world.Seed()
	world.Print()
	for {
		world = world.Step()
		world.Print()

		h = h.Add(world)
		if ContainsDuplicates(h) {
			fmt.Println("Stuck in a loop. Quitting...")
			return
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033c\x0c") // clear terminal
	}
}
