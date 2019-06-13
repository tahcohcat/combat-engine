package entities

import "fmt"

const HAND_SIZE uint8 = 5

type Player struct {

	// user chosen string
	Name string

	// for turn order
	Number uint8

	Characters []Character

	IsAlive bool

	CardSelectChan chan int

}

func NewPlayer(name string, number uint8) Player {
	return Player{
		Name: name,
		Number: number,
		IsAlive: true,
		CardSelectChan : make(chan int),
	}
}

func (player Player) IsPlayerAlive() bool {
	return player.IsAlive
}

func (player Player) PresentMarket(market *Market) {
	fmt.Printf("Player[%s] picking from market:\n%s",
		player.Name, market.Display())
}

func (player *Player) HandleMarketBuy() {
	
}

