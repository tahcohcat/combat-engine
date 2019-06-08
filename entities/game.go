package entities

import "fmt"

type Game struct {
	isRunning bool
}

func (game *Game) start() {
	fmt.Println("Starting rungame...")
}