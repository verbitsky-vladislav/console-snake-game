package base

import (
	"math"
	"snake-game/utils"
)

type DijkstraMovement struct{}

func (d *DijkstraMovement) Move(snake *BaseSnake, foods []*utils.Point, sim *Simulation) {
	if len(foods) == 0 {
		return // Если нет еды, ничего не делаем
	}

	// Найдем ближайшую еду, используя алгоритм Дейкстры
	closestFood := d.findClosestFood(snake, foods)

	// Определим направление к ближайшей еде
	direction := d.determineDirection(snake, closestFood)

	// Проверим, не пытается ли змейка повернуть на 180 градусов
	if !d.isOppositeDirection(snake.Direction, direction) {
		// Двигаем змейку к ближайшей еде
		snake.Direction = direction
	}
}

func (d *DijkstraMovement) findClosestFood(snake *BaseSnake, foods []*utils.Point) *utils.Point {
	if len(foods) == 0 {
		return nil
	}

	var closest *utils.Point
	minDistance := math.Inf(1)

	head := snake.Body[0]

	for _, food := range foods {
		distance := distanceBetweenPoints(head, *food)
		if distance < minDistance {
			minDistance = distance
			closest = food
		}
	}

	return closest
}

func (d *DijkstraMovement) determineDirection(snake *BaseSnake, food *utils.Point) utils.Point {
	head := snake.Body[0]
	foodPosition := *food

	// Вычисляем разницу между головой змеи и позицией еды
	diffX := foodPosition.X - head.X
	diffY := foodPosition.Y - head.Y

	// Выбираем направление на основе разницы
	if diffX > 0 {
		return utils.Point{X: 1, Y: 0} // Двигаемся вправо
	} else if diffX < 0 {
		return utils.Point{X: -1, Y: 0} // Двигаемся влево
	} else if diffY > 0 {
		return utils.Point{X: 0, Y: 1} // Двигаемся вниз
	} else {
		return utils.Point{X: 0, Y: -1} // Двигаемся вверх
	}
}

func (d *DijkstraMovement) isOppositeDirection(current, newDirection utils.Point) bool {
	// Проверяем, являются ли два направления противоположными (различаются на 180 градусов)
	return current.X == -newDirection.X || current.Y == -newDirection.Y
}

func distanceBetweenPoints(p1, p2 utils.Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
