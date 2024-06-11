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
)

func initGame() {
	rand.Seed(time.Now().UnixNano())
	screenWidth, screenHeight = termbox.Size()
	snake = []Point{{screenWidth / 2, screenHeight / 2}}
	food = generateFood()
	direction = Point{0, -1}
}

func generateFood() []Point {
	foodCount := rand.Intn(10) + 1 // Генерируем от 1 до 10 объектов еды
	food := make([]Point, foodCount)
	for i := 0; i < foodCount; i++ {
		food[i] = Point{rand.Intn(screenWidth), rand.Intn(screenHeight)}
	}
	return food
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range snake {
		termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
	for _, f := range food {
		termbox.SetCell(f.x, f.y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
	termbox.Flush()
}

func handleInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			mu.Lock()
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
			mu.Unlock()
		}
	}
}

func update() {
	head := snake[0]
	newHead := Point{head.x + direction.x, head.y + direction.y}

	// Обработка переноса на другую сторону экрана
	if newHead.x >= screenWidth {
		newHead.x = 0
	} else if newHead.x < 0 {
		newHead.x = screenWidth - 1
	}

	if newHead.y >= screenHeight {
		newHead.y = 0
	} else if newHead.y < 0 {
		newHead.y = screenHeight - 1
	}

	for i, f := range food {
		if newHead.x == f.x && newHead.y == f.y {
			food = append(food[:i], food[i+1:]...)                                      // Удалить съеденную еду
			food = append(food, Point{rand.Intn(screenWidth), rand.Intn(screenHeight)}) // Добавить новую еду
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
