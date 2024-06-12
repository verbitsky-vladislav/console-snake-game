package rules

import (
	"math/rand"
	"snake-game/simulation/food"
	"snake-game/simulation/snake"
	"time"
)

func HandleCollision(s1, s2 *snake.BaseSnake, existingFood []*food.Food) {
	if s1.GetHead() == s2.GetHead() {
		s1.Die()
		s2.Die()
		return
	}

	for _, part := range s2.GetBody() {
		if s1.GetHead() == part {
			s1.Die()
			return
		}
	}

	for i, f := range existingFood {
		if f.Position == s1.GetHead() {
			existingFood = append(existingFood[:i], existingFood[i+1:]...)
			s1.Eat(f)
			return
		}
	}
}

func SpawnFood(existingFood []*food.Food, maxFood int, screenWidth, screenHeight int) []*food.Food {
	if len(existingFood) >= maxFood {
		return existingFood
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	newFoodCount := rng.Intn(5) + 1
	for i := 0; i < newFoodCount; i++ {
		x := rng.Intn(screenWidth-2) + 1
		y := rng.Intn(screenHeight-2) + 1
		existingFood = append(existingFood, food.NewFood(x, y))
	}

	return existingFood
}

func ConvertSnakeToFood(snake *snake.BaseSnake) []*food.Food {
	var foodInstance []*food.Food
	for _, part := range snake.GetBody() {
		foodInstance = append(foodInstance, food.NewFood(part.X, part.Y))
	}
	return foodInstance
}
