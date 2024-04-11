package main

import (
	"fmt"
	"main/pkg/common"
	"main/pkg/simulations"
	"time"
)

func main() {
	var g simulations.GameOfLife
	c := g.Start(20, 20)
	h := common.MakeQueue[common.Canvas](3)
	for {
		c.Print()
		c = g.Step(c)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("\033c\x0c") // clear terminal
		h = h.Add(c)
		if g.ShouldStop(h) {
			fmt.Println("Stuck in a loop. Quitting...")
			return
		}
	}
}
