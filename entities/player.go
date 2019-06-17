package entities

import "fmt"

const HAND_SIZE uint8 = 5

type Player struct {

	// user chosen string
	Name string

	// for turn order
	Number uint8

	// Characters held by player
	CharacterHand []Character

	ActionChannel chan Action

	Board PlayerBoard

	IsAlive bool
}

func NewPlayer(name string, number uint8) Player {

	regions := make([]Region, 1)
	regions = append(regions, NewRegion(3))

	return Player{
		Name: name,
		Number: number,
		IsAlive: true,
		ActionChannel : make(chan Action),
		Board : PlayerBoard{
			regions,
		},
	}
}

func (player Player) IsPlayerAlive() bool {
	return player.IsAlive
}

func (player Player) PresentMarket(market *Market) {
	fmt.Printf("Player[%s] picking from market:\n%s",
		player.Name, market.Display())
}

//Callback from client with selection
func (player *Player) HandleMarketBuy() {
	
}