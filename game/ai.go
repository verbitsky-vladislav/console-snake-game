package game

import (
	"math/rand"
	"snake-game/utils"
)

func (s *Snake) UpdateAI(food *FoodManager, player *Snake) {
	if !s.isAlive {
		return
	}

	// Простой ИИ для поиска ближайшей еды
	closestFood := food.GetClosest(s.body[0])
	minDistance := utils.Distance(s.body[0], closestFood)

	// Если змея больше игрока, она пытается его преследовать
	if s.Length() > player.Length() {
		playerHead := player.body[0]
		playerDistance := utils.Distance(s.body[0], playerHead)
		if playerDistance < minDistance {
			closestFood = playerHead
			minDistance = playerDistance
		}
	}

	// Принятие решения о направлении движения
	s.moveToTarget(closestFood)

	// Проверка на столкновение с другими змеями и игроком
	newHead := utils.Point{s.body[0].X + s.direction.X, s.body[0].Y + s.direction.Y}
	if s.checkCollision(newHead, player) || s.checkCollisionWithAISnakes(newHead) {
		s.changeDirection()
		newHead = utils.Point{s.body[0].X + s.direction.X, s.body[0].Y + s.direction.Y}
	}

	s.Update(food)
}

func (s *Snake) checkCollision(newHead utils.Point, other *Snake) bool {
	for _, p := range other.body {
		if p.X == newHead.X && p.Y == newHead.Y {
			return true
		}
	}
	return false
}

func (s *Snake) checkCollisionWithAISnakes(newHead utils.Point) bool {
	for _, aiSnake := range s.game.aiSnakes {
		if aiSnake != s && aiSnake.isAlive && s.checkCollision(newHead, aiSnake) {
			return true
		}
	}
	return false
}

func (s *Snake) changeDirection() {
	// Метод для изменения направления, если возникает столкновение
	directions := []utils.Point{
		{0, 1},  // вниз
		{0, -1}, // вверх
		{1, 0},  // вправо
		{-1, 0}, // влево
	}

	// Пробуем каждое направление, пока не найдем валидное
	for _, dir := range directions {
		// Проверка на противоположное направление
		if dir.X == -s.direction.X && dir.Y == -s.direction.Y {
			continue
		}

		newHead := utils.Point{s.body[0].X + dir.X, s.body[0].Y + dir.Y}
		if !s.checkCollision(newHead, s) && !s.checkCollisionWithAISnakes(newHead) {
			s.direction = dir
			return
		}
	}

	// Если все направления заняты, выбираем случайное направление, чтобы не останавливаться
	s.direction = directions[rand.Intn(len(directions))]
}

func (s *Snake) moveToTarget(target utils.Point) {
	if target.X < s.body[0].X && s.direction.X != 1 {
		s.direction = utils.Point{-1, 0}
	} else if target.X > s.body[0].X && s.direction.X != -1 {
		s.direction = utils.Point{1, 0}
	} else if target.Y < s.body[0].Y && s.direction.Y != 1 {
		s.direction = utils.Point{0, -1}
	} else if target.Y > s.body[0].Y && s.direction.Y != -1 {
		s.direction = utils.Point{0, 1}
	}
}
