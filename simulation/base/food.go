package base

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/utils"
)

type Food []utils.Point

type Manager struct {
	Food         []*utils.Point
	ScreenWidth  int
	ScreenHeight int
	Rng          *rand.Rand
	MaxFood      int
}

func NewFoodManager(screenWidth, screenHeight int, rng *rand.Rand, snakes []*BaseSnake) *Manager {
	return &Manager{
		Food:         generateFood(screenWidth, screenHeight, rng, snakes),
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Rng:          rng,
		MaxFood:      100,
	}
}

func generateFood(screenWidth, screenHeight int, rng *rand.Rand, snakes []*BaseSnake) []*utils.Point {
	foodCount := rng.Intn(10) + 1 // Генерируем от 1 до 10 объектов еды
	food := make([]*utils.Point, 0, foodCount)
	for i := 0; i < foodCount; i++ {
		newFood := generateValidFoodPoint(screenWidth, screenHeight, rng, snakes)
		food = append(food, newFood)
	}
	return food
}

func generateValidFoodPoint(screenWidth, screenHeight int, rng *rand.Rand, snakes []*BaseSnake) *utils.Point {
	for {
		newFood := utils.Point{X: rng.Intn(screenWidth-2) + 1, Y: rng.Intn(screenHeight-2) + 1}
		if !isPositionOccupied(newFood, snakes) {
			return &newFood
		}
	}
}

func isPositionOccupied(point utils.Point, snakes []*BaseSnake) bool {
	for _, snake := range snakes {
		for _, segment := range snake.Body {
			if segment == point {
				return true
			}
		}
	}
	return false
}

func (fm *Manager) Eat(p utils.Point, snakes []*BaseSnake) bool {
	for i, f := range fm.Food {
		if f.X == p.X && f.Y == p.Y {
			fm.Food = append(fm.Food[:i], fm.Food[i+1:]...) // Удалить съеденную еду
			if len(fm.Food) < fm.MaxFood {
				newFoodCount := fm.Rng.Intn(3) + 1 // Добавляем от 1 до 3 объектов еды
				for j := 0; j < newFoodCount && len(fm.Food) < fm.MaxFood; j++ {
					newFood := generateValidFoodPoint(fm.ScreenWidth, fm.ScreenHeight, fm.Rng, snakes)
					fm.Food = append(fm.Food, newFood)
				}
			}
			return true
		}
	}
	return false
}

//func (fm *Manager) GetClosest(p utils.Point) utils.Point {
//	if len(fm.Food) == 0 {
//		return utils.Point{X: -1, Y: -1} // Если нет еды, возвращаем невалидную точку
//	}
//	closestFood := fm.Food[0]
//	minDistance := utils.Distance(p, closestFood)
//	for _, f := range fm.Food {
//		d := utils.Distance(p, f)
//		if d < minDistance {
//			minDistance = d
//			closestFood = f
//		}
//	}
//	return closestFood
//}

func (fm *Manager) Draw() {
	for _, f := range fm.Food {
		termbox.SetCell(f.X, f.Y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
}
