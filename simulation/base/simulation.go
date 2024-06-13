package base

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/utils"
	"sync"
	"time"
)

var snakeColors = []termbox.Attribute{
	termbox.ColorYellow,
	termbox.ColorCyan,
	termbox.ColorMagenta,
	termbox.ColorBlue,
	termbox.ColorWhite,
	termbox.ColorLightGray,
}

type Simulation struct {
	Snakes []*BaseSnake
	Food   *Manager
	Mu     sync.Mutex
}

func NewSimulation(numStartSnakes int) *Simulation {
	simulation := Simulation{}

	screenWidth, screenHeight := termbox.Size()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	snakes := make([]*BaseSnake, 0)

	for i := 0; i < numStartSnakes; i++ {
		newSnake := NewSnake(
			rng.Intn(screenWidth-2)+1,
			rng.Intn(screenHeight-2)+1,
			snakeColors[i%len(snakeColors)],
			//GetRandomAlgorithm(),
			&simulation,
		)
		snakes = append(snakes, newSnake)
	}

	foods := NewFoodManager(screenWidth, screenHeight, rng, snakes)

	simulation.Food = foods
	simulation.Snakes = snakes

	return &simulation
}

func (sim *Simulation) Start() {
	sim.gameLoop()
}

func (sim *Simulation) AddSnake(newSnake *BaseSnake) {
	sim.Snakes = append(sim.Snakes, newSnake)
}

func (sim *Simulation) RemoveSnake(snake *BaseSnake) {
	for i, s := range sim.Snakes {
		if s == snake {
			sim.Snakes = append(sim.Snakes[:i], sim.Snakes[i+1:]...)
			break
		}
	}
}

func (sim *Simulation) gameLoop() {
	ticker := time.NewTicker(50 * time.Millisecond)
	for range ticker.C {
		sim.Mu.Lock()
		sim.update()
		sim.draw()
		sim.Mu.Unlock()
	}
}

func (sim *Simulation) update() {
	for _, snakeInstance := range sim.Snakes {
		if snakeInstance.CheckLife {
			snakeInstance.MovementAlgorithm.Move(snakeInstance, sim.Food.Food, sim)
			snakeInstance.Update(sim.Food)
			sim.checkCollisionsForSnake(snakeInstance)
		} else {
			sim.RemoveSnake(snakeInstance)
		}
	}
}

func (sim *Simulation) checkCollisionsForSnake(snakeInst *BaseSnake) {
	for _, otherSnake := range sim.Snakes {
		if snakeInst == otherSnake || !otherSnake.CheckLife {
			continue
		}

		if snakeInst.CollidesWith(otherSnake) {
			if snakeInst.CollidesHeadWith(otherSnake) {
				snakeInst.Die()
				otherSnake.Die()
			} else if snakeInst.Length() > otherSnake.Length() {
				snakeInst.Eat(otherSnake)
				otherSnake.Die()
			}
		}
	}
}

func (sim *Simulation) draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		utils.LogError.Println("Error clearing screen:", err)
		return
	}

	utils.DrawBorders()
	for _, aiSnake := range sim.Snakes {
		aiSnake.Draw()
	}
	sim.Food.Draw()

	err = termbox.Flush()
	if err != nil {
		utils.LogError.Println("Error flushing screen:", err)
		return
	}
}
