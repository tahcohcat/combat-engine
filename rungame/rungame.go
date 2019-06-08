package main

import "fmt"
import "github.com/tahcohcat/combat-engine/entities"

func main() {

	mechOne := entities.NewGeneticMech()

	fmt.Printf("Mech[%s]:[hp:%03d, ad:%03d, as:%02.2f, m:%03d, c:%03.2f]",
		mechOne.Base.Name,
		mechOne.Base.Stats.Health,
		mechOne.Base.Stats.AttackDamage,
		mechOne.Base.Stats.AttackSpeed,
		mechOne.Base.Stats.MovementSpeed,
		mechOne.Base.Stats.CritPercentage)
}
