package utils

import "github.com/nsf/termbox-go"

func DrawGameOver() {
	screenWidth, screenHeight := termbox.Size()
	message := "Game Over! Press ESC to exit or Enter to restart."
	x := (screenWidth - len(message)) / 2
	y := screenHeight / 2
	for i, ch := range message {
		termbox.SetCell(x+i, y, ch, termbox.ColorRed, termbox.ColorDefault)
	}
}
