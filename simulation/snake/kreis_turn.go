package snake

import (
	"math"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type KreisTurnMovement struct{}

func (kt *KreisTurnMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	path := []utils.Point{}
	angle := 0.0
	step := 1

	for current := start; current != goal; {
		current = utils.Point{
			X: start.X + int(math.Cos(angle)*float64(step)),
			Y: start.Y + int(math.Sin(angle)*float64(step)),
		}
		path = append(path, current)
		angle += math.Pi / 4 // 45 degrees turn
		if angle >= 2*math.Pi {
			angle = 0
			step++
		}
	}

	moveSnakeAlongPath(snake, path)
}
