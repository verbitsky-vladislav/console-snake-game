package snake

import (
	"snake-game/simulation/food"
	"snake-game/utils"
)

type PathfindingMovement struct{}

func (pm *PathfindingMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	// For simplicity, using a basic BFS
	queue := []utils.Point{start}
	cameFrom := make(map[utils.Point]utils.Point)
	cameFrom[start] = start

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == goal {
			break
		}

		for _, neighbor := range getNeighbors(&Node{Position: current}) {
			if _, exists := cameFrom[neighbor]; !exists {
				queue = append(queue, neighbor)
				cameFrom[neighbor] = current
			}
		}
	}

	path := []utils.Point{}
	current := goal
	for current != start {
		path = append([]utils.Point{current}, path...)
		current = cameFrom[current]
	}
	path = append([]utils.Point{start}, path...)

	moveSnakeAlongPath(snake, path)
}
