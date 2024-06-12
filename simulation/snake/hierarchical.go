package snake

import (
	"container/heap"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type HierarchicalMovement struct{}

func (hm *HierarchicalMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	// For simplicity, using a basic A* with larger cells
	cellSize := 5
	startCell := utils.Point{X: start.X / cellSize, Y: start.Y / cellSize}
	goalCell := utils.Point{X: goal.X / cellSize, Y: goal.Y / cellSize}

	openSet := &PriorityQueue{}
	heap.Init(openSet)
	startNode := &Node{
		Position:  startCell,
		Cost:      0,
		Heuristic: heuristic(startCell, goalCell),
		Parent:    nil,
	}
	heap.Push(openSet, startNode)

	cameFrom := make(map[utils.Point]*Node)
	cameFrom[startCell] = startNode

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)
		if current.Position == goalCell {
			path := reconstructPath(current)
			moveSnakeAlongPath(snake, expandPath(path, cellSize))
			return
		}

		for _, neighborPos := range getNeighbors(&Node{Position: current.Position}) {
			newCost := current.Cost + 1
			neighbor, exists := cameFrom[neighborPos]
			if !exists || newCost < neighbor.Cost {
				neighbor = &Node{
					Position:  neighborPos,
					Cost:      newCost,
					Heuristic: heuristic(neighborPos, goalCell),
					Parent:    current,
				}
				cameFrom[neighborPos] = neighbor
				heap.Push(openSet, neighbor)
			}
		}
	}
}

func expandPath(path []utils.Point, cellSize int) []utils.Point {
	expanded := []utils.Point{}
	for _, p := range path {
		expanded = append(expanded, utils.Point{X: p.X * cellSize, Y: p.Y * cellSize})
	}
	return expanded
}
