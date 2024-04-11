package simulations

import (
	"main/pkg/common"
)

type SimulationRunner interface {
	Start(int, int) common.Canvas
	Step() common.Canvas
	ShouldStop() bool
}

type Simulation struct {
	SimulationRunner SimulationRunner
}
