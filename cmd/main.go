package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"snake-game/game"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	// Выбор количества AI-змей
	numAISnakes := selectAISnakes()

	g := game.NewGame(numAISnakes)
	g.Start()
}

func selectAISnakes() int {
	var numAISnakes int
	fmt.Println("Выберите количество AI змей (от 0 до 6):")
	for {
		fmt.Scan(&numAISnakes)
		if numAISnakes >= 0 && numAISnakes <= 6 {
			break
		}
		fmt.Println("Пожалуйста, введите число от 0 до 6.")
	}
	return numAISnakes
}
