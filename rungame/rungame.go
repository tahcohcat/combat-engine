package main

import "fmt"
import "github.com/tahcohcat/combat-engine/entities"

func main() {

	//todo: lobby where players are set up

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

	dragOne := entities.NewGeneticDragon()
	mechOne := entities.NewGeneticMech()

	fmt.Printf("Mech[%s]  : [hp:%03d, ad:%03d, as:%02.2f, m:%03d, c:%03.2f]\n",
		mechOne.Name(),
		mechOne.Stats().Health,
		mechOne.Stats().AttackDamage,
		mechOne.Stats().AttackSpeed,
		mechOne.Stats().MovementSpeed,
		mechOne.Stats().CritPercentage)

	fmt.Printf("Drag[%s]: [hp:%03d, ad:%03d, as:%02.2f, m:%03d, c:%03.2f]\n",
		dragOne.Name(),
		dragOne.Stats().Health,
		dragOne.Stats().AttackDamage,
		dragOne.Stats().AttackSpeed,
		dragOne.Stats().MovementSpeed,
		dragOne.Stats().CritPercentage)
}
