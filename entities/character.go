package entities

type CharacterTribe int

const (
	MechTribe CharacterTribe = 0
)

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

type Character struct {

	Name string
	Tribe CharacterTribe
	Stats Stats
}
