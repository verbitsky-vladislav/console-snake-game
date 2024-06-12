package food

import (
	"github.com/nsf/termbox-go"
	"snake-game/utils"
)

type Food struct {
	Position utils.Point
}

func NewFood(x, y int) *Food {
	return &Food{
		Position: utils.Point{X: x, Y: y},
	}
}

func (f *Food) Draw() {
	termbox.SetCell(f.Position.X, f.Position.Y, 'X', termbox.ColorRed, termbox.ColorDefault)
}
