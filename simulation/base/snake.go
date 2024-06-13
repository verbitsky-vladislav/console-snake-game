package base

import (
	"github.com/nsf/termbox-go"
	"snake-game/utils"
)

type BaseSnake struct {
	Body              []utils.Point
	Direction         utils.Point
	Color             termbox.Attribute
	CheckLife         bool
	MovementAlgorithm Algorithm
	Simulation        *Simulation
}

func NewSnake(x, y int, color termbox.Attribute, sim *Simulation) *BaseSnake {
	return &BaseSnake{
		Body:              []utils.Point{{x, y}},
		Direction:         utils.Point{Y: -1},
		Color:             color,
		CheckLife:         true,
		MovementAlgorithm: &DijkstraMovement{},
		Simulation:        sim,
	}
}

func (s *BaseSnake) Update(food *Manager) {
	if !s.CheckLife {
		utils.LogInfo.Printf("Snake at (%d, %d) is dead and should not update", s.Body[0].X, s.Body[0].Y)
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
			s.Die()
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

	newX, newY := termbox.Size()
	utils.LogInfo.Printf("New snake position: (%d, %d)\n", newX/2, newY/2)

	newSnake := NewSnake(newX/2, newY/2, s.Color, s.Simulation)
	newSnake.Direction = s.Direction

	newSnake.Body = []utils.Point{{X: newX, Y: newY}}

	s.Body = s.Body[:2]

	s.Simulation.AddSnake(newSnake)

	utils.LogInfo.Println("Snake reproduced")
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
	utils.LogInfo.Printf("Snake at (%d, %d) died", s.Body[0].X, s.Body[0].Y)
	s.Simulation.RemoveSnake(s)
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
