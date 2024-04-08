package main

import (
	"fmt"
	"math/rand"
	"time"
	// "example.com/gameoflife"
)

type World [][]bool

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
				fmt.Printf("\xF0\x9F\x9F\xAB")
			default:
				fmt.Printf("\xF0\x9F\x9F\xA9")
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

	// if w.Status(x,y) {
	// 	n--
	// }

	return n
}

func (w World) NextState(x, y int) bool {
	// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
	// Any live cell with two or three live neighbors lives on to the next generation.
	// Any live cell with more than three live neighbors dies, as if by overpopulation.
	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	n := w.Neighbours(x, y)
	// s := w.Status(x,y)
	// return  (n == 3) || (n == 2 && s)
	return n == 3
}

func Step(a, b World) (World, World) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.NextState(i, j)
		}
	}
	return a, b
}

func main() {
	// i := 0
	world := MakeWorld(20, 20)
	nextWorld := MakeWorld(20, 20)
	world.Seed()
	world.Print()
	for {
		nextWorld, world = Step(world, nextWorld)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033c\x0c") // clear terminal
		world.Print()
		// i = i + 1
		// fmt.Println(i)
	}
}
