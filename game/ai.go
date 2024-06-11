package game

import "math/rand"

func (s *Snake) UpdateAI(food *FoodManager, player *Snake) {
	if !s.isAlive {
		return
	}

	// Простой ИИ для поиска ближайшей еды
	closestFood := food.GetClosest(s.body[0])
	minDistance := distance(s.body[0], closestFood)

	// Если змея больше игрока, она пытается его преследовать
	if s.Length() > player.Length() {
		playerHead := player.body[0]
		playerDistance := distance(s.body[0], playerHead)
		if playerDistance < minDistance {
			closestFood = playerHead
			minDistance = playerDistance
		}
	}

	// Принятие решения о направлении движения
	s.moveToTarget(closestFood)

	// Проверка на столкновение с другими змеями и игроком
	newHead := Point{s.body[0].x + s.direction.x, s.body[0].y + s.direction.y}
	if s.checkCollision(newHead, player) || s.checkCollisionWithAISnakes(newHead) {
		s.changeDirection()
		newHead = Point{s.body[0].x + s.direction.x, s.body[0].y + s.direction.y}
	}

	s.Update(food)
}

func (s *Snake) checkCollision(newHead Point, other *Snake) bool {
	for _, p := range other.body {
		if p.x == newHead.x && p.y == newHead.y {
			return true
		}
	}
	return false
}

func (s *Snake) checkCollisionWithAISnakes(newHead Point) bool {
	for _, aiSnake := range s.game.aiSnakes {
		if aiSnake != s && aiSnake.isAlive && s.checkCollision(newHead, aiSnake) {
			return true
		}
	}
	return false
}

func (s *Snake) changeDirection() {
	// Метод для изменения направления, если возникает столкновение
	directions := []Point{
		{0, 1},  // вниз
		{0, -1}, // вверх
		{1, 0},  // вправо
		{-1, 0}, // влево
	}

	// Пробуем каждое направление, пока не найдем валидное
	for _, dir := range directions {
		// Проверка на противоположное направление
		if dir.x == -s.direction.x && dir.y == -s.direction.y {
			continue
		}

		newHead := Point{s.body[0].x + dir.x, s.body[0].y + dir.y}
		if !s.checkCollision(newHead, s) && !s.checkCollisionWithAISnakes(newHead) {
			s.direction = dir
			return
		}
	}

	// Если все направления заняты, выбираем случайное направление, чтобы не останавливаться
	s.direction = directions[rand.Intn(len(directions))]
}

func (s *Snake) moveToTarget(target Point) {
	if target.x < s.body[0].x && s.direction.x != 1 {
		s.direction = Point{-1, 0}
	} else if target.x > s.body[0].x && s.direction.x != -1 {
		s.direction = Point{1, 0}
	} else if target.y < s.body[0].y && s.direction.y != 1 {
		s.direction = Point{0, -1}
	} else if target.y > s.body[0].y && s.direction.y != -1 {
		s.direction = Point{0, 1}
	}
}
