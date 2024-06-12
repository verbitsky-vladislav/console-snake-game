package snake

import (
	"math"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type WavefrontMovement struct{}

func (wm *WavefrontMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	// Initialize wavefront map
	wavefront := make(map[utils.Point]int)
	wavefront[goal] = 0
	queue := []utils.Point{goal}

	// Propagate wavefront
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentWave := wavefront[current]

		for _, neighbor := range getNeighbors(&Node{Position: current}) {
			if _, exists := wavefront[neighbor]; !exists {
				wavefront[neighbor] = currentWave + 1
				queue = append(queue, neighbor)
			}
		}
	}

	// Follow wavefront back to goal
	path := []utils.Point{}
	current := start
	for current != goal {
		path = append(path, current)
		minWave := math.MaxInt
		next := current
		for _, neighbor := range getNeighbors(&Node{Position: current}) {
			if wave, exists := wavefront[neighbor]; exists && wave < minWave {
				minWave = wave
				next = neighbor
			}
		}
		current = next
	}
	path = append(path, goal)

	moveSnakeAlongPath(snake, path)
}
