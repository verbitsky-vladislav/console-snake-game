package base

import "sync"

type Simulation struct {
	Snakes     []*BaseSnake
	Food       *Manager
	Mu         sync.Mutex
	newSnakeCh chan *BaseSnake // Буферизированный канал для добавления новых змей
}
