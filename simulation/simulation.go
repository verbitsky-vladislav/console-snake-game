package base

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/simulation/food"
	"snake-game/simulation/rules"
	"snake-game/simulation/snake"
	"snake-game/utils"
	"time"
)

type Simulation struct {
	Snakes []*snake.BaseSnake
	Food   []*food.Food
}

func NewSimulation() *Simulation {
	return &Simulation{
		Snakes: make([]*snake.BaseSnake, 0),
		Food:   make([]*food.Food, 0),
	}
}

func (s *Simulation) Initialize() {
	screenWidth, screenHeight := utils.GetScreenSize()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		newSnake := snake.NewBaseSnake(screenWidth/2, screenHeight/2)
		s.Snakes = append(s.Snakes, newSnake)
	}

	for i := 0; i < 5; i++ {
		s.Food = append(s.Food, food.NewFood(rng.Intn(screenWidth-2)+1, rng.Intn(screenHeight-2)+1))
	}
}

func (s *Simulation) Run() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		s.Update()
		s.Draw()
	}
}

func (s *Simulation) Update() {
	for _, sn := range s.Snakes {
		if sn.IsAlive() {
			sn.Move(s.Food)
		}
	}

	aliveSnakes := s.Snakes[:0]
	for _, sn := range s.Snakes {
		if sn.IsAlive() {
			aliveSnakes = append(aliveSnakes, sn)
		} else {
			// Преобразование мертвой змейки в еду и добавление в список еды
			foodFromSnake := rules.ConvertSnakeToFood(sn)
			s.Food = append(s.Food, foodFromSnake...)
		}
	}
	s.Snakes = aliveSnakes

	screenWidth, screenHeight := utils.GetScreenSize()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(s.Food); i < 5; i++ {
		newFoodCount := rng.Intn(5) + 1
		for j := 0; j < newFoodCount; j++ {
			s.Food = append(s.Food, food.NewFood(rng.Intn(screenWidth-2)+1, rng.Intn(screenHeight-2)+1))
		}
	}
}

func (s *Simulation) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	utils.DrawBorders()
	for _, sn := range s.Snakes {
		sn.Draw()
	}
	for _, fd := range s.Food {
		fd.Draw()
	}
	termbox.Flush()
}
