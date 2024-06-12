package snake

import (
	"math/rand"
	"snake-game/simulation/food"
)

type Algorithm interface {
	Move(snake *BaseSnake, food []*food.Food)
}

type AlgorithmType int

const (
	AStar AlgorithmType = iota
	Dijkstra
	Wavefront
	Pathfinding
	Navmesh
	Hierarchical
	ObstacleAvoidance
	DivideAndConquer
	KreisTurn
)

var MovementAlgorithms = []Algorithm{
	&AStarMovement{},
	//&DijkstraMovement{},
	//&WavefrontMovement{},
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
