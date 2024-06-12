package snake

import (
	"container/heap"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type DijkstraMovement struct{}

func (dm *DijkstraMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	openSet := &PriorityQueue{}
	heap.Init(openSet)
	startNode := &Node{
		Position:  start,
		Cost:      0,
		Heuristic: 0,
		Parent:    nil,
	}
	heap.Push(openSet, startNode)

	cameFrom := make(map[utils.Point]*Node)
	cameFrom[start] = startNode

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)
		if current.Position == goal {
			path := reconstructPath(current)
			moveSnakeAlongPath(snake, path)
			return
		}

		for _, neighborPos := range getNeighbors(current) {
			newCost := current.Cost + 1
			neighbor, exists := cameFrom[neighborPos]
			if !exists || newCost < neighbor.Cost {
				neighbor = &Node{
					Position:  neighborPos,
					Cost:      newCost,
					Heuristic: 0,
					Parent:    current,
				}
				cameFrom[neighborPos] = neighbor
				heap.Push(openSet, neighbor)
			}
		}
	}
}