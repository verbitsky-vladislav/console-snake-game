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

type Snake struct {
	body      []Point
	direction Point
	isAlive   bool
}

var (
	player       Snake
	aiSnakes     []Snake
	food         []Point
	screenWidth  int
	screenHeight int
	mu           sync.Mutex
)

func initGame() {
	rand.Seed(time.Now().UnixNano())
	screenWidth, screenHeight = termbox.Size()
	player = Snake{body: []Point{{screenWidth / 2, screenHeight / 2}}, direction: Point{0, -1}, isAlive: true}
	aiSnakes = []Snake{
		{body: []Point{{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}}, direction: Point{0, 1}, isAlive: true},
		{body: []Point{{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}}, direction: Point{1, 0}, isAlive: true},
	}
	food = generateFood()
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
	drawSnake(player, termbox.ColorGreen)
	for _, snake := range aiSnakes {
		drawSnake(snake, termbox.ColorYellow)
	}
	for _, f := range food {
		termbox.SetCell(f.x, f.y, 'X', termbox.ColorRed, termbox.ColorDefault)
	}
	if !player.isAlive {
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

func drawSnake(snake Snake, color termbox.Attribute) {
	for _, p := range snake.body {
		termbox.SetCell(p.x, p.y, 'O', color, termbox.ColorDefault)
	}
}

func drawGameOver() {
	message := "Game Over! Press ESC to exit or Enter to restart."
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
			if !player.isAlive {
				if ev.Key == termbox.KeyEsc {
					termbox.Close()
					return
				}
				if ev.Key == termbox.KeyEnter {
					initGame()
				}
			} else {
				switch ev.Key {
				case termbox.KeyArrowUp:
					if player.direction.y == 0 {
						player.direction = Point{0, -1}
					}
				case termbox.KeyArrowDown:
					if player.direction.y == 0 {
						player.direction = Point{0, 1}
					}
				case termbox.KeyArrowLeft:
					if player.direction.x == 0 {
						player.direction = Point{-1, 0}
					}
				case termbox.KeyArrowRight:
					if player.direction.x == 0 {
						player.direction = Point{1, 0}
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
	if !player.isAlive {
		return
	}

	updateSnake(&player)
	for i := range aiSnakes {
		updateAISnake(&aiSnakes[i])
	}

	checkCollisions()
}

func updateSnake(snake *Snake) {
	head := snake.body[0]
	newHead := Point{head.x + snake.direction.x, head.y + snake.direction.y}

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
	for _, p := range snake.body {
		if p.x == newHead.x && p.y == newHead.y {
			snake.isAlive = false
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
	snake.body = snake.body[:len(snake.body)-1] // Удалить хвост, если еда не съедена
SkipReduction:
	snake.body = append([]Point{newHead}, snake.body...)
}

func updateAISnake(snake *Snake) {
	if !snake.isAlive {
		return
	}

	// Простой ИИ для поиска ближайшей еды
	closestFood := food[0]
	minDistance := distance(snake.body[0], closestFood)
	for _, f := range food {
		d := distance(snake.body[0], f)
		if d < minDistance {
			minDistance = d
			closestFood = f
		}
	}

	if closestFood.x < snake.body[0].x {
		snake.direction = Point{-1, 0}
	} else if closestFood.x > snake.body[0].x {
		snake.direction = Point{1, 0}
	} else if closestFood.y < snake.body[0].y {
		snake.direction = Point{0, -1}
	} else if closestFood.y > snake.body[0].y {
		snake.direction = Point{0, 1}
	}

	updateSnake(snake)
}

func distance(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}

func checkCollisions() {
	for i := range aiSnakes {
		for _, s := range aiSnakes {
			if !s.isAlive {
				continue
			}
			if len(player.body) > len(s.body) && snakeCollision(player.body[0], s.body) {
				s.isAlive = false
				player.body = append(player.body, s.body...)
				respawnAISnake(&s)
			} else if len(s.body) > len(player.body) && snakeCollision(s.body[0], player.body) {
				player.isAlive = false
				s.body = append(s.body, player.body...)
				initGame()
				return
			}
		}

		for j := range aiSnakes {
			if i == j || !aiSnakes[j].isAlive {
				continue
			}
			if len(aiSnakes[i].body) > len(aiSnakes[j].body) && snakeCollision(aiSnakes[i].body[0], aiSnakes[j].body) {
				aiSnakes[j].isAlive = false
				aiSnakes[i].body = append(aiSnakes[i].body, aiSnakes[j].body...)
				respawnAISnake(&aiSnakes[j])
			}
		}
	}
}

func snakeCollision(head Point, body []Point) bool {
	for _, p := range body {
		if head.x == p.x && head.y == p.y {
			return true
		}
	}
	return false
}

func respawnAISnake(snake *Snake) {
	snake.body = []Point{{rand.Intn(screenWidth-2) + 1, rand.Intn(screenHeight-2) + 1}}
	snake.direction = Point{0, 1}
	snake.isAlive = true
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
