package game

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

type Game struct {
	player   *Snake
	aiSnakes []*Snake
	food     *FoodManager
	mu       sync.Mutex
}

func NewGame(numAISnakes int) *Game {
	screenWidth, screenHeight := termbox.Size()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	game := &Game{}
	player := NewSnake(screenWidth/2, screenHeight/2, termbox.ColorGreen, game, true) // true для игрока
	aiSnakes := make([]*Snake, numAISnakes)
	for i := 0; i < numAISnakes; i++ {
		aiSnakes[i] = NewSnake(rng.Intn(screenWidth-2)+1, rng.Intn(screenHeight-2)+1, snakeColors[i], game, false)
	}
	food := NewFoodManager(screenWidth, screenHeight, rng)

	game.player = player
	game.aiSnakes = aiSnakes
	game.food = food

	return game
}

func (g *Game) Start() {
	go g.handleInput()
	g.gameLoop()
}

func (g *Game) handleInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			g.mu.Lock()
			if !g.player.isAlive {
				if ev.Key == termbox.KeyEsc {
					termbox.Close()
					return
				}
				if ev.Key == termbox.KeyEnter {
					g.player.RespawnWithLength(1)
				}
			} else {
				switch ev.Key {
				case termbox.KeyArrowUp:
					if g.player.direction.Y == 0 {
						g.player.direction = utils.Point{0, -1}
					}
				case termbox.KeyArrowDown:
					if g.player.direction.Y == 0 {
						g.player.direction = utils.Point{0, 1}
					}
				case termbox.KeyArrowLeft:
					if g.player.direction.X == 0 {
						g.player.direction = utils.Point{-1, 0}
					}
				case termbox.KeyArrowRight:
					if g.player.direction.X == 0 {
						g.player.direction = utils.Point{1, 0}
					}
				case termbox.KeyCtrlC:
					termbox.Close()
					return
				}
			}
			g.mu.Unlock()
		}
	}
}

func (g *Game) gameLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		g.mu.Lock()
		g.update()
		g.draw()
		g.mu.Unlock()
	}
}

func (g *Game) update() {
	g.player.Update(g.food)
	for _, aiSnake := range g.aiSnakes {
		if aiSnake.isAlive {
			aiSnake.UpdateAI(g.food, g.player)
		} else {
			aiSnake.RespawnWithLength(1)
		}
	}

	g.checkCollisions()
}

func (g *Game) checkCollisions() {
	// Проверка столкновений игрока с AI-змеями
	for _, aiSnake := range g.aiSnakes {
		if !aiSnake.isAlive {
			continue
		}
		if g.player.CollidesWith(aiSnake) {
			if g.player.CollidesHeadWith(aiSnake) {
				g.player.TurnBodyToFood(g.food)
				g.player.Die()
				aiSnake.TurnBodyToFood(g.food)
				aiSnake.Die()
			} else {
				aiSnake.Eat(g.player)
				g.player.RespawnWithLength(1) // Убедимся, что этот метод вызывается
			}
		}

		for _, otherAISnake := range g.aiSnakes {
			if aiSnake != otherAISnake && otherAISnake.isAlive {
				if aiSnake.CollidesWith(otherAISnake) {
					if aiSnake.CollidesHeadWith(otherAISnake) {
						aiSnake.TurnBodyToFood(g.food)
						aiSnake.Die()
						otherAISnake.TurnBodyToFood(g.food)
						otherAISnake.Die()
					} else if aiSnake.Length() > otherAISnake.Length() {
						aiSnake.Eat(otherAISnake)
						otherAISnake.RespawnWithLength(1)
					}
				}
			}
		}
	}
}

func (g *Game) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	utils.DrawBorders()
	g.player.Draw()
	for _, aiSnake := range g.aiSnakes {
		aiSnake.Draw()
	}
	g.food.Draw()
	termbox.Flush()
}
