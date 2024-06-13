package base

import (
	"math/rand"
	"snake-game/utils"
)

type Manager struct {
	Food         []*utils.Point
	ScreenWidth  int
	ScreenHeight int
	Rng          *rand.Rand
	MaxFood      int
}
