package snake

import (
	"container/heap"
	"math"
	"snake-game/simulation/food"
	"snake-game/utils"
)

type AStarMovement struct{}

type Node struct {
	Position  utils.Point
	Cost      int
	Heuristic int
	Parent    *Node
	Index     int // Index of the item in the heap.
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost+pq[i].Heuristic < pq[j].Cost+pq[j].Heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

func heuristic(a, b utils.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func getNeighbors(node *Node) []utils.Point {
	return []utils.Point{
		{X: node.Position.X + 1, Y: node.Position.Y},
		{X: node.Position.X - 1, Y: node.Position.Y},
		{X: node.Position.X, Y: node.Position.Y + 1},
		{X: node.Position.X, Y: node.Position.Y - 1},
	}
}

func (am *AStarMovement) Move(snake *BaseSnake, food []*food.Food) {
	if len(food) == 0 {
		return
	}

	start := snake.GetHead()
	goal := food[0].Position

	openSet := &PriorityQueue{}
	heap.Init(openSet)
	startNode := &Node{
		Position:  start,
		Cost:      0,
		Heuristic: heuristic(start, goal),
		Parent:    nil,
	}
	heap.Push(openSet, startNode)

	cameFrom := make(map[utils.Point]*Node)
	cameFrom[start] = startNode

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)
		if current.Position == goal {
			path := reconstructPath(current)
			moveSnakeAlongPath(snake, path)
			return
		}

		for _, neighborPos := range getNeighbors(current) {
			newCost := current.Cost + 1
			neighbor, exists := cameFrom[neighborPos]
			if !exists || newCost < neighbor.Cost {
				neighbor = &Node{
					Position:  neighborPos,
					Cost:      newCost,
					Heuristic: heuristic(neighborPos, goal),
					Parent:    current,
				}
				cameFrom[neighborPos] = neighbor
				heap.Push(openSet, neighbor)
			}
		}
	}
}

func reconstructPath(node *Node) []utils.Point {
	path := []utils.Point{}
	for node != nil {
		path = append(path, node.Position)
		node = node.Parent
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func moveSnakeAlongPath(snake *BaseSnake, path []utils.Point) {
	if len(path) > 1 {
		next := path[1]
		// Переместить голову змейки на следующую позицию
		snake.Body[0] = next
		// Установить направление на основе разницы между текущей головой и следующей позицией
		direction := utils.Point{X: next.X - snake.GetHead().X, Y: next.Y - snake.GetHead().Y}
		// Инвертируем направление Y для соответствия координатной системе termbox
		direction.Y = -direction.Y
		snake.SetDirection(direction)
	}
}
