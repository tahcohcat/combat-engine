package entities

import "fmt"

type Game struct {
	IsRunning bool

	Players []Player

	RoundNumber uint
	MaxRounds uint
}

func (game Game) numberPlayers() int {
	return len(Filter(game.Players, Player.IsPlayerAlive))
}
func (game *Game) StartNewRound() {

	game.RoundNumber++

	for _, player := range game.Players {
		market := Market{size: uint8(game.numberPlayers()) }
		market.Populate()
		player.TradeOnMarket(&market)
	}

}

func (game *Game) Start() {
	fmt.Println("Starting run-game...")

	for !game.isDone() {
		alivePlayers := Filter(game.Players, Player.IsPlayerAlive)
		fmt.Printf("Round[%d] : [%d/%d] players alive\n",
			game.RoundNumber,
			len(alivePlayers),
			len(game.Players))

		game.StartNewRound()
	}
}

func (game *Game) isDone() bool {
	return game.numberPlayers() == 0 ||
		game.RoundNumber == game.MaxRounds
}