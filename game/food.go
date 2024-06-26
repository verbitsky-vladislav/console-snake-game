package game

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/utils"
)

type FoodManager struct {
	food         []utils.Point
	screenWidth  int
	screenHeight int
	rng          *rand.Rand
	maxFood      int
}

func NewFoodManager(screenWidth, screenHeight int, rng *rand.Rand) *FoodManager {
	return &FoodManager{
		food:         generateFood(screenWidth, screenHeight, rng, nil, nil),
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		rng:          rng,
		maxFood:      20,
	}
}

func generateFood(screenWidth, screenHeight int, rng *rand.Rand, snakes []*Snake, player *Snake) []utils.Point {
	foodCount := rng.Intn(10) + 1 // Генерируем от 1 до 10 объектов еды
	food := make([]utils.Point, 0, foodCount)
	for i := 0; i < foodCount; i++ {
		newFood := generateValidFoodPoint(screenWidth, screenHeight, rng, snakes, player)
		food = append(food, newFood)
	}
	return food
}

func generateValidFoodPoint(screenWidth, screenHeight int, rng *rand.Rand, snakes []*Snake, player *Snake) utils.Point {
	for {
		newFood := utils.Point{rng.Intn(screenWidth-2) + 1, rng.Intn(screenHeight-2) + 1}
		if !isPositionOccupied(newFood, snakes, player) {
			return newFood
		}
	}
}

func isPositionOccupied(point utils.Point, snakes []*Snake, player *Snake) bool {
	if player != nil {
		for _, segment := range player.body {
			if segment == point {
				return true
			}
		}
	}
	for _, snake := range snakes {
		for _, segment := range snake.body {
			if segment == point {
				return true
			}
		}
	}
	return false
}

func (fm *FoodManager) Eat(p utils.Point, snakes []*Snake, player *Snake) bool {
	for i, f := range fm.food {
		if f.X == p.X && f.Y == p.Y {
			fm.food = append(fm.food[:i], fm.food[i+1:]...) // Удалить съеденную еду
			if len(fm.food) < fm.maxFood {
				newFoodCount := fm.rng.Intn(3) + 1 // Добавляем от 1 до 3 объектов еды
				for j := 0; j < newFoodCount && len(fm.food) < fm.maxFood; j++ {
					newFood := generateValidFoodPoint(fm.screenWidth, fm.screenHeight, fm.rng, snakes, player)
					fm.food = append(fm.food, newFood)
				}
			}
			return true
		}
	}
	return false
}

func (fm *FoodManager) GetClosest(p utils.Point) utils.Point {
	if len(fm.food) == 0 {
		return utils.Point{-1, -1} // Если нет еды, возвращаем невалидную точку
	}
	closestFood := fm.food[0]
	minDistance := utils.Distance(p, closestFood)
	for _, f := range fm.food {
		d := utils.Distance(p, f)
		if d < minDistance {
			minDistance = d
			closestFood = f
		}
	}
	return closestFood
}

func (fm *FoodManager) Draw() {
	for _, f := range fm.food {
		termbox.SetCell(f.X, f.Y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
}
