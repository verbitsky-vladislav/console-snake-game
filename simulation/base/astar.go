package base

//
//import (
//	"container/heap"
//	"github.com/nsf/termbox-go"
//	_ "math"
//	"snake-game/utils"
//)
//
//type AStarMovement struct{}
//
//type node struct {
//	point utils.Point
//	fCost float64
//	gCost float64
//	parent *node
//}
//
//func (a *AStarMovement) Move(snake *BaseSnake, foods []*utils.Point) {
//	if len(foods) == 0 {
//		return // Если нет еды, ничего не делаем
//	}
//
//	head := snake.Body[0]
//	grid := createGrid(snake, foods)
//
//	// Запускаем алгоритм A*
//	path := aStarAlgorithm(grid, head, foods[0])
//
//	// Находим направление к следующему шагу на пути к еде
//	if len(path) > 1 {
//		nextStep := path[len(path)-2].point
//		direction := utils.Point{X: nextStep.X - head.X, Y: nextStep.Y - head.Y}
//
//		// Проверяем, не пытается ли змейка повернуть на 180 градусов
//		if !a.isOppositeDirection(snake.Direction, direction) {
//			// Двигаем змейку к следующему шагу на пути
//			snake.Direction = direction
//		}
//	}
//}
//
//func createGrid(snake *BaseSnake, foods []*utils.Point) [][]int {
//	screenWidth, screenHeight := termbox.Size()
//	grid := make([][]int, screenHeight)
//	for i := range grid {
//		grid[i] = make([]int, screenWidth)
//	}
//
//	// Заполняем сетку препятствиями (тело змейки и другие змейки)
//	for _, snakeInstance := range sim.Snakes {
//		for _, segment := range snakeInstance.Body {
//			grid[segment.Y][segment.X] = -1 // -1 означает препятствие
//		}
//	}
//
//	// Помечаем позиции еды на сетке
//	for _, food := range foods {
//		grid[food.Y][food.X] = 1 // 1 означает позицию еды
//	}
//
//	return grid
//}
//
//func aStarAlgorithm(grid [][]int, start, end utils.Point) []*node {
//	openSet := priorityQueue{}
//	heap.Init(&openSet)
//
//	startNode := &node{point: start, fCost: 0, gCost: 0, parent: nil}
//	heap.Push(&openSet, startNode)
//
//	closedSet := make(map[utils.Point]bool)
//
//	var current *node
//
//	for openSet.Len() > 0 {
//		current = heap.Pop(&openSet).(*node)
//
//		if current.point == end {
//			break
//		}
//
//		closedSet[current.point] = true
//
//		for _, neighbor := range neighbors(current.point, grid) {
//			if closedSet[neighbor.point] {
//				continue
//			}
//
//			// Расстояние от начала до соседа
//			gCost := current.gCost + distanceBetweenPoints(current.point, neighbor.point)
//
//			if !openSetContains(openSet, neighbor) || gCost < neighbor.gCost {
//				neighbor.gCost = gCost
//				neighbor.fCost = gCost + heuristicCost(neighbor.point, end)
//				neighbor.parent = current
//
//				if !openSetContains(openSet, neighbor) {
//					heap.Push(&openSet, neighbor)
//				}
//			}
//		}
//	}
//
//	// Строим путь от конечной точки к начальной
//	path := []*node{}
//	for current != nil {
//		path = append(path, current)
//		current = current.parent
//	}
//
//	// Разворачиваем путь, чтобы начало пути было в начале списка
//	for i := 0; i < len(path)/2; i++ {
//		j := len(path) - i - 1
//		path[i], path[j] = path[j], path[i]
//	}
//
