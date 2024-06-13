package utils

type Point struct {
	X, Y int
}

func Distance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ClampInt(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
