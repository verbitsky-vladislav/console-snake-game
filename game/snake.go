package game

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

type Snake struct {
	body      []Point
	direction Point
	color     termbox.Attribute
	isAlive   bool
	game      *Game // Добавляем ссылку на игру
	isPlayer  bool  // Указываем, является ли змея игроком
}

func NewSnake(x, y int, color termbox.Attribute, game *Game, isPlayer bool) *Snake {
	return &Snake{
		body:      []Point{{x, y}},
		direction: Point{0, -1},
		color:     color,
		isAlive:   true,
		game:      game, // Инициализируем ссылку на игру
		isPlayer:  isPlayer,
	}
}

func (s *Snake) Update(food *FoodManager) {
	if !s.isAlive {
		return
	}

	head := s.body[0]
	newHead := Point{head.x + s.direction.x, head.y + s.direction.y}

	// Обработка переноса на другую сторону экрана
	screenWidth, screenHeight := termbox.Size()
	if newHead.x >= screenWidth-1 {
		newHead.x = 1
	} else if newHead.x < 1 {
		newHead.x = screenWidth - 2
	}

	if newHead.y >= screenHeight-1 {
		newHead.y = 1
	} else if newHead.y < 1 {
		newHead.y = screenHeight - 2
	}

	// Проверка на столкновение с собой
	for _, p := range s.body {
		if p.x == newHead.x && p.y == newHead.y {
			s.TurnBodyToFood(food)
			s.isAlive = false
			if s.isPlayer {
				s.RespawnWithLength(1) // Перезапускаем игрока с длиной 1
			}
			return
		}
	}

	if food.Eat(newHead, s.game.aiSnakes, s.game.player) {
		s.body = append([]Point{newHead}, s.body...)
	} else {
		s.body = append([]Point{newHead}, s.body...)
		s.body = s.body[:len(s.body)-1]
	}
}

func (s *Snake) CollidesWith(other *Snake) bool {
	for _, p := range other.body {
		if s.body[0].x == p.x && s.body[0].y == p.y {
			return true
		}
	}
	return false
}

func (s *Snake) CollidesHeadWith(other *Snake) bool {
	return s.body[0] == other.body[0]
}

func (s *Snake) Eat(other *Snake) {
	s.body = append(s.body, other.body...)
}

func (s *Snake) Die() {
	s.isAlive = false
}

func (s *Snake) TurnBodyToFood(food *FoodManager) {
	for _, segment := range s.body {
		food.food = append(food.food, segment)
	}
}

func (s *Snake) Respawn() {
	screenWidth, screenHeight := termbox.Size()
	s.body = []Point{{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}}
	s.direction = Point{0, 1}
	s.isAlive = true
}

func (s *Snake) RespawnWithLength(length int) {
	screenWidth, screenHeight := termbox.Size()
	s.body = make([]Point, length)
	for i := 0; i < length; i++ {
		s.body[i] = Point{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}
	}
	s.direction = Point{0, 1}
	s.isAlive = true
}

func (s *Snake) Draw() {
	for _, p := range s.body {
		termbox.SetCell(p.x, p.y, 'O', s.color, termbox.ColorDefault)
	}
}

func (s *Snake) Length() int {
	return len(s.body)
}
