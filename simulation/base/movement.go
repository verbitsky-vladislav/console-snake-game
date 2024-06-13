package base

import (
	"math/rand"
	"snake-game/utils"
)

type Algorithm interface {
	Move(snake *BaseSnake, food []*utils.Point)
}

type AlgorithmType int

const (
	//AStar AlgorithmType = iota
	Dijkstra AlgorithmType = iota
	//Wavefront
	//Pathfinding
	//Navmesh
	//Hierarchical
	//ObstacleAvoidance
	//DivideAndConquer
	//KreisTurn
)

var MovementAlgorithms = []Algorithm{
	//&AStarMovement{},
	&DijkstraMovement{},
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
