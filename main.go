package main

import (
	"fmt"
	"main/pkg/common"
	"main/pkg/simulations"
	"time"
)

func main() {
	c := simulations.Start(20, 20)
	h := common.MakeQueue[common.Canvas](3)
	for {
		c.Print()
		c = simulations.Step(c)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("\033c\x0c") // clear terminal
		h = h.Add(c)
		if simulations.ShouldStop(h) {
			fmt.Println("Stuck in a loop. Quitting...")
			return
		}
	}
}
