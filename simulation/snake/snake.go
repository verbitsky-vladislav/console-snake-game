package snake

import (
	"github.com/nsf/termbox-go"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type BaseSnakeInterface interface {
	Move(food []*food.Food)
	Eat(food *food.Food)
	Reproduce() BaseSnakeInterface
	Die()
	IsAlive() bool
	GetBody() []utils.Point
	SetDirection(dir utils.Point)
	GetDirection() utils.Point
	GetHead() utils.Point
}

type BaseSnake struct {
	Body      []utils.Point
	Direction utils.Point
	CheckLife bool
	Algorithm Algorithm
}

func NewBaseSnake(x, y int) *BaseSnake {
	return &BaseSnake{
		Body:      []utils.Point{{X: x, Y: y}},
		Direction: utils.Point{X: 0, Y: -1},
		CheckLife: true,
		Algorithm: GetRandomAlgorithm(),
	}
}

func (s *BaseSnake) Move(food []*food.Food) {
	if s.CheckLife {
		s.Algorithm.Move(s, food)
	}
}

func (s *BaseSnake) Eat(food *food.Food) {
	s.Body = append([]utils.Point{food.Position}, s.Body...)
}

func (s *BaseSnake) Reproduce() *BaseSnake {
	if len(s.Body) >= 5 {
		s.Body = s.Body[:2]
		return NewBaseSnake(s.Body[0].X, s.Body[0].Y)
	}
	return nil
}

func (s *BaseSnake) Die() {
	s.CheckLife = false
}

func (s *BaseSnake) IsAlive() bool {
	return s.CheckLife
}

func (s *BaseSnake) GetBody() []utils.Point {
	return s.Body
}

func (s *BaseSnake) SetDirection(dir utils.Point) {
	s.Direction = dir
}

func (s *BaseSnake) GetDirection() utils.Point {
	return s.Direction
}

func (s *BaseSnake) GetHead() utils.Point {
	return s.Body[0]
}

func (s *BaseSnake) Draw() {
	for _, p := range s.Body {
		termbox.SetCell(p.X, p.Y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
}
