package gameoflife

import (
	"fmt"
	"math/rand"
)

type World [][]bool

func MakeWorld(height int, width int) World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}

	seed(w)
	return w
}

func seed(w World) {
	for _, row := range w {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func Print(w World) {
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
