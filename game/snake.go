package game

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/utils"
)

type Snake struct {
	body      []utils.Point
	direction utils.Point
	color     termbox.Attribute
	isAlive   bool
	game      *Game // Добавляем ссылку на игру
	isPlayer  bool  // Указываем, является ли змея игроком
}

func NewSnake(x, y int, color termbox.Attribute, game *Game, isPlayer bool) *Snake {
	return &Snake{
		body:      []utils.Point{{x, y}},
		direction: utils.Point{Y: -1},
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
	newHead := utils.Point{head.X + s.direction.X, head.Y + s.direction.Y}

	// Обработка переноса на другую сторону экрана
	screenWidth, screenHeight := termbox.Size()
	if newHead.X >= screenWidth-1 {
		newHead.X = 1
	} else if newHead.X < 1 {
		newHead.X = screenWidth - 2
	}

	if newHead.Y >= screenHeight-1 {
		newHead.Y = 1
	} else if newHead.Y < 1 {
		newHead.Y = screenHeight - 2
	}

	// Проверка на столкновение с собой
	for _, p := range s.body {
		if p.X == newHead.X && p.Y == newHead.Y {
			s.TurnBodyToFood(food)
			s.isAlive = false
			if s.isPlayer {
				s.RespawnWithLength(1) // Перезапускаем игрока с длиной 1
			}
			return
		}
	}

	if food.Eat(newHead, s.game.aiSnakes, s.game.player) {
		s.body = append([]utils.Point{newHead}, s.body...)
	} else {
		s.body = append([]utils.Point{newHead}, s.body...)
		s.body = s.body[:len(s.body)-1]
	}
}

func (s *Snake) CollidesWith(other *Snake) bool {
	for _, p := range other.body {
		if s.body[0].X == p.X && s.body[0].Y == p.Y {
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
	s.body = []utils.Point{{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}}
	s.direction = utils.Point{0, 1}
	s.isAlive = true
}

func (s *Snake) RespawnWithLength(length int) {
	screenWidth, screenHeight := termbox.Size()
	s.body = make([]utils.Point, length)
	for i := 0; i < length; i++ {
		s.body[i] = utils.Point{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}
	}
	s.direction = utils.Point{0, 1}
	s.isAlive = true
}

func (s *Snake) Draw() {
	for _, p := range s.body {
		termbox.SetCell(p.X, p.Y, 'O', s.color, termbox.ColorDefault)
	}
}

func (s *Snake) Length() int {
	return len(s.body)
}
