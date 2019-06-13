package entities

import (
	"fmt"
	"log"
	"math/rand"
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

func (game *Game) start(
	Player1Select chan int,
	Player2Select chan int,
	GameEnd chan int) {

	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			log.Print("End of round\n")
		case <- Player1Select:
			log.Printf("Player 1 has selected a card [%d]", <- Player1Select)
		case <- Player2Select:
			log.Printf("Player 2 has selected a card [%d]", <- Player2Select)
		case <-GameEnd:
			log.Print("End of game")
			return
		}
	}

}

func (game *Game) WaitForInput(p *Player, wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("go-routine: [wait for input from player.%d: %s] started\n",
		p.Number, p.Name)

	for {

		//s := rand.Intn(3 - 1) + 1
		//time.Sleep(time.Duration(s) * time.Second)

		time.Sleep(3 * time.Second)
		cardSelection := rand.Intn(8 - 1) + 1

		log.Printf("Player [%s] selecting card [%d]...", p.Name, cardSelection)

		p.CardSelectChan <- cardSelection
	}
}

func (game *Game) Start() {

	var wg sync.WaitGroup

	fmt.Println("Starting run-game...")

	GameEnd := make(chan int)

	wg.Add(1)
	go game.WaitForInput(&game.Players[0], &wg)

	wg.Add(1)
	go game.WaitForInput(&game.Players[1], &wg)

	log.Println("Game.start...")

	game.start(
		game.Players[0].CardSelectChan,
		game.Players[1].CardSelectChan,
		GameEnd)

	wg.Wait()

	log.Println("End of game?")
}

func (game *Game) isDone() bool {

	alivePlayers := Filter(game.Players, Player.IsPlayerAlive)
	fmt.Printf("gamae.is-done: [%d/%d] players alive\n",
		len(alivePlayers), len(game.Players))

	return game.numberPlayers() == 0 ||
		game.RoundNumber == game.MaxRounds
}