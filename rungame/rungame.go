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

	playerOne := entities.Player{
		Name: "player-one",
		Number: 0,
		Characters: make([]entities.Character, 0),
		IsAlive : true}

	playerTwo := entities.Player{
		Name: "player-two",
		Number: 1,
		Characters: make([]entities.Character, 0),
		IsAlive : true}

	players = append(players, playerOne)
	players = append(players, playerTwo)

	game := entities.Game{
		IsRunning : false,
		Players : players,
		RoundNumber: 0,
		MaxRounds: 10 }

	game.Start()

	log.Println("Game has ended. Waiting for HTTP server shutdown")

	waitGroup.Wait()
}
