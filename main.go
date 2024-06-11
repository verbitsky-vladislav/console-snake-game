package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"sync"
	"time"
)

type Point struct {
	x, y int
}

var (
	snake        []Point
	food         []Point
	direction    Point
	screenWidth  int
	screenHeight int
	mu           sync.Mutex
	gameOver     bool
)

func initGame() {
	rand.Seed(time.Now().UnixNano())
	screenWidth, screenHeight = termbox.Size()
	snake = []Point{{screenWidth / 2, screenHeight / 2}}
	food = generateFood()
	direction = Point{0, -1}
	gameOver = false
}

func generateFood() []Point {
	foodCount := rand.Intn(10) + 1 // Генерируем от 1 до 10 объектов еды
	food := make([]Point, foodCount)
	for i := 0; i < foodCount; i++ {
		food[i] = Point{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}
	}
	return food
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBorders()
	for _, p := range snake {
		termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
	for _, f := range food {
		termbox.SetCell(f.x, f.y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
	if gameOver {
		drawGameOver()
	}
	termbox.Flush()
}

func drawBorders() {
	for x := 0; x < screenWidth; x++ {
		termbox.SetCell(x, 0, '#', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(x, screenHeight-1, '#', termbox.ColorWhite, termbox.ColorDefault)
	}
	for y := 0; y < screenHeight; y++ {
		termbox.SetCell(0, y, '#', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(screenWidth-1, y, '#', termbox.ColorWhite, termbox.ColorDefault)
	}
}

func drawGameOver() {
	message := "Game Over! Press ESC to exit."
	x := (screenWidth - len(message)) / 2
	y := screenHeight / 2
	for i, ch := range message {
		termbox.SetCell(x+i, y, ch, termbox.ColorRed, termbox.ColorDefault)
	}
}

func handleInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			mu.Lock()
			if gameOver {
				if ev.Key == termbox.KeyEsc {
					termbox.Close()
					return
				}
			} else {
				switch ev.Key {
				case termbox.KeyArrowUp:
					if direction.y == 0 {
						direction = Point{0, -1}
					}
				case termbox.KeyArrowDown:
					if direction.y == 0 {
						direction = Point{0, 1}
					}
				case termbox.KeyArrowLeft:
					if direction.x == 0 {
						direction = Point{-1, 0}
					}
				case termbox.KeyArrowRight:
					if direction.x == 0 {
						direction = Point{1, 0}
					}
				case termbox.KeyCtrlC:
					termbox.Close()
					return
				}
			}
			mu.Unlock()
		}
	}
}

func update() {
	if gameOver {
		return
	}

	head := snake[0]
	newHead := Point{head.x + direction.x, head.y + direction.y}

	// Обработка переноса на другую сторону экрана
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
	for _, p := range snake {
		if p.x == newHead.x && p.y == newHead.y {
			gameOver = true
			return
		}
	}

	for i, f := range food {
		if newHead.x == f.x && newHead.y == f.y {
			food = append(food[:i], food[i+1:]...)                                                  // Удалить съеденную еду
			food = append(food, Point{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}) // Добавить новую еду
			goto SkipReduction
		}
	}
	snake = snake[:len(snake)-1] // Удалить хвост, если еда не съедена
SkipReduction:
	snake = append([]Point{newHead}, snake...)
}

func gameLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for range ticker.C {
		mu.Lock()
		update()
		draw()
		mu.Unlock()
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	initGame()
	go handleInput()
	gameLoop()
}
