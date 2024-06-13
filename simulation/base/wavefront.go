package base

//
//import (
//	"github.com/nsf/termbox-go"
//	"math"
//	"snake-game/utils"
//)
//
//type WaveMovement struct{}
//
//func (w *WaveMovement) Move(snake *BaseSnake, foods []*utils.Point, sim *Simulation) {
//	if len(foods) == 0 {
//		return // Если нет еды, ничего не делаем
//	}
//
//	head := snake.Body[0]
//	grid := createGrid(snake, foods, sim)
//
//	// Запускаем волновой алгоритм
//	waveGrid := waveAlgorithm(grid, head)
//
//	// Находим путь к ближайшей еде
//	direction := findPathToFood(waveGrid, head, snake.Direction)
//
//	// Проверяем, не пытается ли змейка повернуть на 180 градусов
//	if !w.isOppositeDirection(snake.Direction, direction) {
//		// Двигаем змейку к ближайшей еде
//		snake.Direction = direction
//	}
//}
//
//func createGrid(snake *BaseSnake, foods []*utils.Point, sim *Simulation) [][]int {
//	screenWidth, screenHeight := termbox.Size()
//	grid := make([][]int, screenHeight)
//	for i := range grid {
//		grid[i] = make([]int, screenWidth)
//	}
//
//	// Заполняем сетку препятствиями (тело змейки и другие змейки)
//	for _, s := range sim.Snakes {
//		for _, segment := range s.Body {
//			if segment.Y >= 0 && segment.Y < len(grid) && segment.X >= 0 && segment.X < len(grid[0]) {
//				grid[segment.Y][segment.X] = -1 // -1 означает препятствие
//			}
//		}
//	}
//
//	// Помечаем позиции еды на сетке
//	for _, food := range foods {
//		if food.Y >= 0 && food.Y < len(grid) && food.X >= 0 && food.X < len(grid[0]) {
//			grid[food.Y][food.X] = 1 // 1 означает позицию еды
//		}
//	}
//
//	return grid
//}
//
//func waveAlgorithm(grid [][]int, start utils.Point) [][]int {
//	queue := []utils.Point{start}
//	directions := []utils.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Возможные направления движения: влево, вправо, вверх, вниз
//
//	waveGrid := make([][]int, len(grid))
//	for i := range grid {
//		waveGrid[i] = make([]int, len(grid[0]))
//		copy(waveGrid[i], grid[i])
//	}
//
//	for len(queue) > 0 {
//		current := queue[0]
//		queue = queue[1:]
//
//		for _, dir := range directions {
//			next := utils.Point{X: current.X + dir.X, Y: current.Y + dir.Y}
//
//			if isValid(next, grid) && waveGrid[next.Y][next.X] == 0 {
//				waveGrid[next.Y][next.X] = waveGrid[current.Y][current.X] + 1
//				queue = append(queue, next)
//			}
//		}
//	}
//
//	return waveGrid
//}
//
//func isValid(point utils.Point, grid [][]int) bool {
//	if point.Y < 0 || point.Y >= len(grid) || point.X < 0 || point.X >= len(grid[0]) {
//		return false
//	}
//	return true
//}
//
//func findPathToFood(waveGrid [][]int, start utils.Point, currentDirection utils.Point) utils.Point {
//	directions := []utils.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Возможные направления движения: влево, вправо, вверх, вниз
//
//	minDistance := math.Inf(1)
//	var bestDirection utils.Point
//
//	for _, dir := range directions {
//		next := utils.Point{X: start.X + dir.X, Y: start.Y + dir.Y}
//
//		if isValid(next, waveGrid) && waveGrid[next.Y][next.X] > 0 {
//			distance := float64(waveGrid[next.Y][next.X])
//
//			// Предпочтительнее двигаться в направлении текущего движения
//			if dir == currentDirection {
//				distance -= 0.1 // Чтобы учесть предпочтения змейки
//			}
//
//			if distance < minDistance {
//				minDistance = distance
//				bestDirection = dir
//			}
//		}
//	}
//
//	// Если не удалось найти направление, возвращаем текущее направление змейки
//	if bestDirection == (utils.Point{}) {
//		return currentDirection
//	}
//
//	return bestDirection
//}
//
//func (w *WaveMovement) isOppositeDirection(current, newDirection utils.Point) bool {
//	return current.X == -newDirection.X || current.Y == -newDirection.Y
//}
