package utils

import "github.com/nsf/termbox-go"

func GetScreenSize() (int, int) {
	return termbox.Size()
}

func DrawBorders() {
	width, height := termbox.Size()
	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, height-1, '-', termbox.ColorDefault, termbox.ColorDefault)
	}
	for y := 0; y < height; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(width-1, y, '|', termbox.ColorDefault, termbox.ColorDefault)
	}
}
