package game

type Point struct {
	x, y int
}

func distance(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}
