package entities

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Game struct {
	IsRunning bool

	Players []Player

	RoundNumber uint
	MaxRounds uint

	Id uint
}

func (game Game) numberPlayers() int {
	return len(Filter(game.Players, Player.IsPlayerAlive))
}


func (game *Game) playAnotherRound() {

	var waitGroup sync.WaitGroup

	game.RoundNumber++

	timerCallback := make(chan bool, 1)
	waitGroup.Add(1)
	go SetTimer(&waitGroup, timerCallback, 5 * time.Second)
	var isTimerDone bool = false

	//round - phase 1 - markets
	for _, player := range game.Players {
		market := Market{size: uint8(game.numberPlayers()) + 1}
		market.Populate()
		player.PresentMarket(&market)
	}
	log.Println("All players received their cards")

	isTimerDone = <-timerCallback
	log.Printf("Timer return as %t\n", isTimerDone)

}

func (game *Game) Start() {

	fmt.Println("Starting run-game...")

	for !game.isDone() {
		game.playAnotherRound()
	}
}

func (game *Game) isDone() bool {

	alivePlayers := Filter(game.Players, Player.IsPlayerAlive)
	fmt.Printf("gamae.is-done: [%d/%d] players alive\n",
		len(alivePlayers), len(game.Players))

	return game.numberPlayers() == 0 ||
		game.RoundNumber == game.MaxRounds
}