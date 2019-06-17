package entities

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type RoundPhase int

const (
	MarketPhase RoundPhase = 0
	CharPlacePhase RoundPhase = 1
	FightPhase RoundPhase = 2
	RoundEndPhase RoundPhase = 3
)

type Game struct {
	IsRunning bool

	Players []Player

	Markets []Market

	CurrentRoundPhase RoundPhase
	RoundNumber uint
	MaxRounds uint

	Id uint
}

func (game Game) numberPlayers() int {
	return len(Filter(game.Players, Player.IsPlayerAlive))
}

func (game *Game) start(Player1Action chan Action, Player2Action chan Action, GameEnd chan int) {

	//market-phase
	log.Println("Setting up markets")
	game.SetUpCurrentPhase()

	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			log.Printf("End of %s-phase for round %d \n", RoundPhaseAsString(game.CurrentRoundPhase), game.RoundNumber)
			game.CurrentRoundPhase = game.CurrentRoundPhase +  1
		case <-Player1Action:
			action := <-Player1Action
			if game.IsValidAction(action) {
				log.Printf("Player 1 [%s]", action.AsString())
			} else {
				log.Printf("Player 1 [%s] was ILLEGAL in %s",
					action.AsString(), RoundPhaseAsString(game.CurrentRoundPhase))
			}
		case <-Player2Action:
			action := <-Player2Action
			if game.IsValidAction(action) {
				log.Printf("Player 2 [%s]", action.AsString())
			} else {
				log.Printf("Player 2 [%s] was ILLEGAL in %s",
					action.AsString(), RoundPhaseAsString(game.CurrentRoundPhase))
			}
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

		time.Sleep(3 * time.Second)

		if game.CurrentRoundPhase == MarketPhase {
			// simulating player performing appropriate action
			market := game.Markets[p.Number]

			pickedCharacterIndex := randomInt16(0, int16(len(market.CharacterOptions)))
			resource := market.CharacterOptions[pickedCharacterIndex].Resource()

			log.Printf("Player [%s] picking card [%d]... [%s | %s | type:%d]",
				p.Name, pickedCharacterIndex, resource.Name, resource.Id.String(), resource.Type)
			p.ActionChannel <- NewPickCharacterAction(resource)
		}
		if game.CurrentRoundPhase == CharPlacePhase {

			pickedCharacterIndex := randomInt16(0, int16(len(p.CharacterHand)))
			resource := p.CharacterHand[pickedCharacterIndex].Resource()

			log.Printf("Player [%s] playing card [%d]... [%s | %s | type:%d]",
				p.Name, pickedCharacterIndex, resource.Name, resource.Id.String(), resource.Type)

			p.ActionChannel <- NewPickCharacterAction(resource)
		}
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
		game.Players[0].ActionChannel,
		game.Players[1].ActionChannel,
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

func (game *Game) SetUpCurrentPhase() {
	log.Printf("Setting up game-round-phase %s\n", RoundPhaseAsString(game.CurrentRoundPhase))

	switch game.CurrentRoundPhase {
	case MarketPhase: {
		for index := range game.Players {
			game.Markets[index].Populate()
		}
	}
	default: {
		panic("Implement me")
	}
	}
}
func (game Game) IsValidAction(action Action) bool {
	switch game.CurrentRoundPhase {
	case MarketPhase: return action.ActionType() == PickCharacter
	case CharPlacePhase: return action.ActionType() == PlayCharacter
	default: return false
	}
}