package entities

import "fmt"

type Dragon struct {
	base CharacterBase
}

func (d Dragon) Display() string {
	return d.base.Display()
}

func (d Dragon) Name() string {
	return d.base.Name
}

func (d Dragon) Stats() Stats {
	return d.base.Stats
}

func (Dragon) Tribe() CharacterTribe {
	return DragonTribe
}

func NewDragon() Character {
	c  := Dragon{}
	c.base.Name = "DragonName"
	c.base.Tribe = DragonTribe
	c.base.Stats = dragonBaseStats()
	return c
}

func dragonBaseStats() Stats {
	return Stats {
		Health : 75,
		AttackDamage : 70,
		AttackSpeed : 0.8,
		MovementSpeed : 60,
		CritPercentage : 0.0,
	}
}

func NewGeneticDragon() Character {

	character  := NewDragon()

	dragon, ok := character.(Dragon)
	if !ok {
		fmt.Errorf("We expected a dragon, but got type %T \n", character)
		fmt.Printf("We expected a dragon, but got type %T \n", character)
		panic("We expected a dragon")
	}

	dragon.base.Stats.Health += uint16(randomInt16(-10, 10))
	dragon.base.Stats.AttackDamage += uint16(randomInt16(-10, 10))
	dragon.base.Stats.AttackSpeed += randomFloat32(-0.3, 0.3)
	dragon.base.Stats.MovementSpeed += uint16(randomInt16(-10, 10))
	return &dragon
}
