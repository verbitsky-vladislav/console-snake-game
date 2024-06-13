package base

import (
	"github.com/nsf/termbox-go"
	"snake-game/utils"
)

type BaseSnake struct {
	Body              []utils.Point
	Direction         utils.Point
	Color             termbox.Attribute
	CheckLife         bool
	MovementAlgorithm Algorithm
	Simulation        *Simulation
}
