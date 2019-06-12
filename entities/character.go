package entities

import "fmt"

type CharacterTribe int
type ResourceType int

const NumTribes = 2

const (
	MechTribe CharacterTribe = 0
	DragonTribe CharacterTribe = 1
)

const (
	ResGold ResourceType = 0
	ResSpecial ResourceType = 1
	ResLight ResourceType = 2
	ResDark ResourceType = 3
)

type Cost struct {

	//resource types list
	//and their quantities
	Resources [] map[ResourceType]int
}

type Stats struct {
	Health uint16
	AttackDamage uint16
	AttackSpeed float32
	MovementSpeed uint16
	CritPercentage float32
}

func (s *Stats) setDefaults() {
	s.Health = 100
	s.AttackDamage = 50
	s.AttackSpeed = 1.0
	s.MovementSpeed = 40
	s.CritPercentage = 0.0
}

type Character interface {
	Name() string
	Stats() Stats
	Tribe() CharacterTribe
	Display() string
}

type CharacterBase struct {
	Name string
	Tribe CharacterTribe
	Stats Stats
	Cost Cost
}

func (b CharacterBase) Display() string {
	return fmt.Sprintf("[%s] [hp:%03d, ad:%03d, as:%02.2f, m:%03d, c:%03.2f]\n",
		tribeToString(b.Tribe),
		b.Stats.Health,
		b.Stats.AttackDamage,
		b.Stats.AttackSpeed,
		b.Stats.MovementSpeed,
		b.Stats.CritPercentage)
}