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
	Snakes     []*BaseSnake
	Food       *Manager
	Mu         sync.Mutex
	newSnakeCh chan *BaseSnake // Буферизированный канал для добавления новых змей
}

func NewSimulation(numStartSnakes int) *Simulation {
	simulation := Simulation{}

	screenWidth, screenHeight := termbox.Size()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	snakes := make([]*BaseSnake, 0) // Инициализируем слайс для хранения змеек

	for i := 0; i < numStartSnakes; i++ {
		// Создаем новую змейку и добавляем ее в слайс
		newSnake := NewSnake(
			rng.Intn(screenWidth-2)+1,
			rng.Intn(screenHeight-2)+1,
			snakeColors[i%len(snakeColors)], // корректируем индекс, чтобы не выйти за пределы массива цветов
			GetRandomAlgorithm(),
			&simulation,
		)
		snakes = append(snakes, newSnake)
	}

	foods := NewFoodManager(screenWidth, screenHeight, rng, snakes)

	// Возвращаем новый экземпляр симуляции
	simulation.Food = foods
	simulation.Snakes = snakes
	simulation.newSnakeCh = make(chan *BaseSnake, 100) // Увеличиваем размер буфера канала

	return &simulation
}

func (sim *Simulation) Start() {
	utils.LogInfo.Println("Simulation started")
	sim.gameLoop()
}

func (sim *Simulation) AddSnake(newSnake *BaseSnake) {
	utils.LogInfo.Println("Attempting to add a new snake")

	sim.Mu.Lock()
	defer sim.Mu.Unlock()

	// Запускаем обновление новой змеи в отдельной горутине перед добавлением в список
	go func() {
		for range time.Tick(100 * time.Millisecond) {
			sim.Mu.Lock()
			newSnake.Update(sim.Food)
			sim.Mu.Unlock()
		}
	}()

	// Добавляем новую змею к симуляции напрямую
	sim.Snakes = append(sim.Snakes, newSnake)

	utils.LogInfo.Println("New snake added to list")
}

func (sim *Simulation) gameLoop() {
	ticker := time.NewTicker(50 * time.Millisecond)
	for range ticker.C {
		sim.Mu.Lock()
		sim.update()
		sim.draw()
		sim.Mu.Unlock()

		// Обработка канала для добавления новых змей
		select {
		case newSnake := <-sim.newSnakeCh:
			utils.LogInfo.Println("New snake added from channel to list")
			sim.Snakes = append(sim.Snakes, newSnake)
		default:
		}
	}
}

func (sim *Simulation) update() {
	utils.LogInfo.Println("Updating simulation")

	for _, snakeInst := range sim.Snakes {
		if snakeInst.CheckLife {
			snakeInst.MovementAlgorithm.Move(snakeInst, sim.Food.Food)
			snakeInst.Update(sim.Food)
		} else {
			snakeInst.Die()
		}
	}

	sim.checkCollisions()
}

func (sim *Simulation) checkCollisions() {
	utils.LogInfo.Println("Checking collisions")

	for _, snakeInst := range sim.Snakes {
		if !snakeInst.CheckLife {
			continue
		}

		for _, otherSnake := range sim.Snakes {
			if snakeInst != otherSnake && otherSnake.CheckLife {
				if snakeInst.CollidesWith(otherSnake) {
					if snakeInst.CollidesHeadWith(otherSnake) {
						snakeInst.TurnBodyToFood(sim.Food)
						snakeInst.Die()
						otherSnake.TurnBodyToFood(sim.Food)
						otherSnake.Die()
					} else if snakeInst.Length() > otherSnake.Length() {
						snakeInst.Eat(otherSnake)
						otherSnake.Die()
					}
				}
			}
		}
	}
}

func (sim *Simulation) draw() {
	utils.LogInfo.Println("Drawing simulation")

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
