package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"snake-game/game"
	base "snake-game/simulation"
	"snake-game/ui"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	mode := selectMode()

	if mode == 1 {
		numAISnakes := selectAISnakes()
		g := game.NewGame(numAISnakes)
		go ui.HandleInput(g)
		g.Start()
	} else if mode == 2 {
		s := base.NewSimulation()
		s.Initialize()
		go ui.HandleInput(s)
		s.Run()
	} else {
		fmt.Println("Invalid mode selected")
	}
}

func selectMode() uint8 {
	var mode uint8
	fmt.Println("Выберите режим: game (1) или simulation (2)")
	for {
		fmt.Scan(&mode)
		if mode == 1 || mode == 2 {
			break
		}
		fmt.Println("Пожалуйста, введите game или simulation.")
	}
	return mode
}

func selectAISnakes() int {
	var numAISnakes int
	fmt.Println("Выберите количество AI змей (от 0 до 5):")
	for {
		fmt.Scan(&numAISnakes)
		if numAISnakes >= 0 && numAISnakes <= 5 {
			break
		}
		fmt.Println("Пожалуйста, введите число от 0 до 5.")
	}
	return numAISnakes
}
