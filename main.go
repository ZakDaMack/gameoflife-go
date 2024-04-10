package main

import (
	"fmt"
	"main/gameoflife"
	"time"
)

func ContainsDuplicates(q gameoflife.Queue[gameoflife.World]) bool {
	for i := 0; i < q.Length()-1; i++ {
		for j := i + 1; j < q.Length(); j++ {
			if q.Get(i).IsSame(q.Get(j)) {
				return true
			}
		}
	}

	return false
}

func main() {
	// i := 0
	world := gameoflife.MakeWorld(20, 20)
	h := gameoflife.MakeQueue[gameoflife.World](3)
	world.Seed()
	for {
		world.Print()
		time.Sleep(500 * time.Millisecond)
		world = world.Step()
		fmt.Println("\033c\x0c") // clear terminal

		h = h.Add(world)
		if ContainsDuplicates(h) {
			fmt.Println("Stuck in a loop. Quitting...")
			return
		}
	}
}
