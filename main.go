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
	food         Point
	direction    Point
	screenWidth  int
	screenHeight int
	mu           sync.Mutex
)

func initGame() {
	rand.Seed(time.Now().UnixNano())
	screenWidth, screenHeight = termbox.Size()
	snake = []Point{{screenWidth / 2, screenHeight / 2}}
	food = Point{rand.Intn(screenWidth), rand.Intn(screenHeight)}
	direction = Point{0, -1}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range snake {
		termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)
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
			case termbox.KeyEsc:
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

	if newHead.x == food.x && newHead.y == food.y {
		food = Point{rand.Intn(screenWidth), rand.Intn(screenHeight)}
	} else {
		snake = snake[:len(snake)-1]
	}
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
