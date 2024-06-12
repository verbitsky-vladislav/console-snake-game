package ui

import (
	"github.com/nsf/termbox-go"
)

func HandleInput(g interface{}) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				termbox.Close()
				return
			}
		}
	}
}
