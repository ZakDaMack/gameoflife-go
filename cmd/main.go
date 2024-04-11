package main

import (
	"flag"
	"fmt"
	"main/pkg/simulations"
	"os"
	"time"
)

const (
	timeMs = 500
)

func main() {
	game := os.Args[len(os.Args)-1]
	seed := flag.String("seed", "", "Seed used to generate the starting position")
	h := flag.Int("height", 20, "Canvas height")
	w := flag.Int("width", 20, "Canvas width")
	flag.Parse()

	var s simulations.SimulationRunner
	switch game {
	case "gameoflife":
		s = simulations.MakeGameOfLife(*seed)
	default:
		fmt.Println("Simulation", game, "does not exist")
		os.Exit(1)
	}

	c := s.Start(*h, *w)
	for {
		c.Print()
		c = s.Step()
		time.Sleep(timeMs * time.Millisecond)
		fmt.Println("\033c\x0c") // clear terminal
		if s.ShouldStop() {
			fmt.Println("Stuck in a loop. Quitting...")
			return
		}
	}
}
