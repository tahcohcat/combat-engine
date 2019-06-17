package main

import (
	"github.com/tahcohcat/combat-engine/entities"
	"log"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {

	/*
	HTTP SERVER - TEST CODE
	waitGroup.Add(1)
	go rest.StartHTTPServer(8000, &waitGroup)

	//wait for server startup
	time.Sleep(1)
	*/

	players := make([]entities.Player, 0)
	players = append(players, entities.NewPlayer("player-one", 0))
	players = append(players, entities.NewPlayer("player-two", 1))

	markets := make([]entities.Market, 0)
	markets = append(markets, entities.Market{ CharacterOptions: make([]entities.Character, 5), Size: 5})
	markets = append(markets, entities.Market{ CharacterOptions: make([]entities.Character, 5), Size: 5})

	game := entities.Game{
		IsRunning : false,
		Players : players,
		CurrentRoundPhase: entities.MarketPhase,
		Markets : markets,
		RoundNumber: 0,
		MaxRounds: 10 }

	game.Start()

	log.Println("Game has ended. Waiting for HTTP server shutdown")

	waitGroup.Wait()
}
