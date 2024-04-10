package simulations

import (
	"main/pkg/common"
)

type Simulation interface {
	Start() common.Canvas
	Step() common.Canvas
	ShouldStop() bool
}
