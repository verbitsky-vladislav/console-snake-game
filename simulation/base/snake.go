package base

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/utils"
)

type BaseSnake struct {
	//base.BaseSnake
	Body              []utils.Point
	Direction         utils.Point
	Color             termbox.Attribute
	CheckLife         bool
	MovementAlgorithm Algorithm
	Simulation        *Simulation
}

func NewSnake(x, y int, color termbox.Attribute, algorithm Algorithm, sim *Simulation) *BaseSnake {
	return &BaseSnake{
		Body:              []utils.Point{{x, y}},
		Direction:         utils.Point{Y: -1},
		Color:             color,
		CheckLife:         true,
		MovementAlgorithm: algorithm,
		Simulation:        sim,
	}
}

func (s *BaseSnake) Update(food *Manager) {
	if !s.CheckLife {
		return
	}

	head := s.Body[0]
	newHead := utils.Point{X: head.X + s.Direction.X, Y: head.Y + s.Direction.Y}

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
	for _, p := range s.Body {
		if p.X == newHead.X && p.Y == newHead.Y {
			s.TurnBodyToFood(food)
			s.CheckLife = false
			return
		}
	}
	if food.Eat(newHead, s.Simulation.Snakes) {
		s.Body = append([]utils.Point{newHead}, s.Body...)
	} else {
		s.Body = append([]utils.Point{newHead}, s.Body...)
		s.Body = s.Body[:len(s.Body)-1]
	}

	// Проверка на размножение
	if len(s.Body) == 5 {
		s.Reproduce()
	}
}

func (s *BaseSnake) Reproduce() {
	utils.LogInfo.Println("Reproducing snake")

	//newX, newY := s.calculateReproductionPosition()
	screenWidth, screenHeight := termbox.Size()
	utils.LogInfo.Printf("New snake position: (%d, %d)\n", screenWidth/2, screenHeight/2)

	newSnake := NewSnake(screenWidth/2, screenHeight/2, s.Color, s.MovementAlgorithm, s.Simulation)
	newSnake.Direction = s.Direction
	newSnake.Body = make([]utils.Point, len(s.Body))
	copy(newSnake.Body, s.Body)
	s.Simulation.AddSnake(newSnake)

	// Уменьшаем длину родительской змейки
	s.Body = s.Body[:2]

	utils.LogInfo.Println("Snake reproduced")
}

func (s *BaseSnake) calculateReproductionPosition() (int, int) {
	utils.LogInfo.Println("Calculating reproduction position")

	s.Simulation.Mu.Lock()
	defer s.Simulation.Mu.Unlock()

	screenWidth, screenHeight := termbox.Size()
	newX, newY := s.Body[0].X, s.Body[0].Y

	// Определяем новую позицию за пределами области, занимаемой родительской змеей
	for i := 0; i < 10; i++ { // Пытаемся найти новую позицию не более 10 раз
		newX = utils.ClampInt(newX+rand.Intn(3)-1, 1, screenWidth-2)
		newY = utils.ClampInt(newY+rand.Intn(3)-1, 1, screenHeight-2)
		if !s.positionOccupied(newX, newY) { // Проверяем, не занята ли новая позиция другой змеёй
			break
		}
	}

	utils.LogInfo.Printf("New position: (%d, %d)\n", newX, newY)

	// Устанавливаем случайное направление для новой змейки
	directions := []utils.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	newDirection := directions[rand.Intn(len(directions))]

	// Применяем новое направление к новой змейке
	s.Direction = newDirection

	return newX, newY
}

func (s *BaseSnake) positionOccupied(x, y int) bool {
	for _, snake := range s.Simulation.Snakes {
		for _, point := range snake.Body {
			if point.X == x && point.Y == y {
				return true
			}
		}
	}
	return false
}

func (s *BaseSnake) CollidesWith(other *BaseSnake) bool {
	for _, p := range other.Body {
		if s.Body[0].X == p.X && s.Body[0].Y == p.Y {
			return true
		}
	}
	return false
}

func (s *BaseSnake) CollidesHeadWith(other *BaseSnake) bool {
	return s.Body[0] == other.Body[0]
}

func (s *BaseSnake) Eat(other *BaseSnake) {
	s.Body = append(s.Body, other.Body...)
}

func (s *BaseSnake) Die() {
	s.CheckLife = false
}

func (s *BaseSnake) TurnBodyToFood(food *Manager) {
	for _, segment := range s.Body {
		food.Food = append(food.Food, &segment)
	}
}

func (s *BaseSnake) Draw() {
	for _, p := range s.Body {
		termbox.SetCell(p.X, p.Y, 'O', s.Color, termbox.ColorDefault)
	}
}

func (s *BaseSnake) Length() int {
	return len(s.Body)
}
