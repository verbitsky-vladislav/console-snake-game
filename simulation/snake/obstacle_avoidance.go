package snake

import (
	"math"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type ObstacleAvoidanceMovement struct{}

func (oa *ObstacleAvoidanceMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	// Basic potential field
	avoidance := utils.Point{X: 0, Y: 0}
	for _, obstacle := range snake.Body {
		direction := utils.Point{X: start.X - obstacle.X, Y: start.Y - obstacle.Y}
		distance := math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y))
		if distance == 0 {
			continue
		}
		influence := 1 / distance
		avoidance.X += int(float64(direction.X) * influence)
		avoidance.Y += int(float64(direction.Y) * influence)
	}

	next := utils.Point{
		X: start.X + goal.X - start.X + avoidance.X,
		Y: start.Y + goal.Y - start.Y + avoidance.Y,
	}
	snake.SetDirection(utils.Point{X: next.X - start.X, Y: next.Y - start.Y})
	snake.Body = append([]utils.Point{next}, snake.Body...)
}
