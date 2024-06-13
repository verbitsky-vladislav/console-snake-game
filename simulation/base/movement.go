package base

import (
	"math/rand"
	"snake-game/utils"
)

type Algorithm interface {
	Move(snake *BaseSnake, food []*utils.Point, simulation *Simulation)
}

var MovementAlgorithms = []Algorithm{
	//&AStarMovement{},
	&DijkstraMovement{},
	//&WaveMovement{},
	//&PathfindingMovement{},
	//&NavmeshMovement{},
	//&HierarchicalMovement{},
	//&ObstacleAvoidanceMovement{},
	//&DivideAndConquerMovement{},
	//&KreisTurnMovement{},
}

func GetRandomAlgorithm() Algorithm {
	return MovementAlgorithms[rand.Intn(len(MovementAlgorithms))]
}
