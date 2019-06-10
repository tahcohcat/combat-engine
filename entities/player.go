package entities

import "fmt"

type Player struct {

	// user chosen string
	Name string

	// for turn order
	Number uint8

	Characters []Character

	IsAlive bool
}

func (player Player) IsPlayerAlive() bool {
	return player.IsAlive
}

func (player Player) TradeOnMarket(market *Market) {
	fmt.Printf("Player[%s] picking from market:\n%s",
		player.Name, market.Display())
}
