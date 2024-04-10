package common

import (
	"fmt"
)

type Canvas [][]bool

func MakeCanvas(height, width int) Canvas {
	w := make(Canvas, height)
	for i := range w {
		w[i] = make([]bool, width)
	}

	return w
}

func (c Canvas) Size() (int, int) {
	if len(c) == 0 {
		return 0, 0
	}
	// returns height, width
	return len(c), len(c[0])
}

func (c Canvas) Print() {
	for _, row := range c {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf("\xF0\x9F\x9F\xA9")
				// fmt.Print("X")
			default:
				fmt.Printf("\xF0\x9F\x9F\xAB")
				// fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (c Canvas) Copy() Canvas {
	newWorld := MakeCanvas(c.Size())
	for i := range c {
		copy(newWorld[i], c[i])
	}
	return newWorld
}

func (c Canvas) IsSame(canvas Canvas) bool {
	for i := range c {
		for j := range c[i] {
			if canvas[i][j] != c[i][j] {
				return false
			}
		}
	}
	return true
}

func (c Canvas) GetCell(x, y int) bool {
	height, width := c.Size()
	b := (height + y) % height
	a := (width + x) % width
	return c[b][a]
}
