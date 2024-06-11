package game

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

type FoodManager struct {
	food         []Point
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

func generateFood(screenWidth, screenHeight int, rng *rand.Rand, snakes []*Snake, player *Snake) []Point {
	foodCount := rng.Intn(10) + 1 // Генерируем от 1 до 10 объектов еды
	food := make([]Point, 0, foodCount)
	for i := 0; i < foodCount; i++ {
		newFood := generateValidFoodPoint(screenWidth, screenHeight, rng, snakes, player)
		food = append(food, newFood)
	}
	return food
}

func generateValidFoodPoint(screenWidth, screenHeight int, rng *rand.Rand, snakes []*Snake, player *Snake) Point {
	for {
		newFood := Point{rng.Intn(screenWidth-2) + 1, rng.Intn(screenHeight-2) + 1}
		if !isPositionOccupied(newFood, snakes, player) {
			return newFood
		}
	}
}

func isPositionOccupied(point Point, snakes []*Snake, player *Snake) bool {
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

func (fm *FoodManager) Eat(p Point, snakes []*Snake, player *Snake) bool {
	for i, f := range fm.food {
		if f.x == p.x && f.y == p.y {
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

func (fm *FoodManager) GetClosest(p Point) Point {
	if len(fm.food) == 0 {
		return Point{-1, -1} // Если нет еды, возвращаем невалидную точку
	}
	closestFood := fm.food[0]
	minDistance := distance(p, closestFood)
	for _, f := range fm.food {
		d := distance(p, f)
		if d < minDistance {
			minDistance = d
			closestFood = f
		}
	}
	return closestFood
}

func (fm *FoodManager) Draw() {
	for _, f := range fm.food {
		termbox.SetCell(f.x, f.y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
}
