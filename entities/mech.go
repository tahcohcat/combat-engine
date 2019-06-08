package entities

import (
	"math/rand"
	"time"
)

type Mech struct {
	Base Character
}

func NewStandardMech() Mech {
	m  := Mech{}

	m.Base.Name = "MechName"
	m.Base.Tribe = MechTribe

	m.Base.Stats.Health = 100
	m.Base.Stats.AttackDamage = 50
	m.Base.Stats.AttackSpeed = 1.0
	m.Base.Stats.MovementSpeed = 40
	m.Base.Stats.CritPercentage = 0.0

	return m
}

func NewGeneticMech() Mech {

	m  := Mech{}

	m.Base.Name = "MechName"
	m.Base.Tribe = MechTribe

	m.Base.Stats.Health = uint16(100 + randomInt16(-10, 10))
	m.Base.Stats.AttackDamage = uint16(20 + randomInt16(-10, 10))
	m.Base.Stats.AttackSpeed = 1.0 + randomFloat32(-0.3, 0.3)
	m.Base.Stats.MovementSpeed = uint16(100 + randomInt16(-10, 10))
	m.Base.Stats.CritPercentage = 0.0

	return m
}

func randomFloat32(min, max float32) float32 {
	rand.Seed(time.Now().Unix())
	return min + rand.Float32() * (max - min)
}

func randomInt16(min, max int16) int16 {
	rand.Seed(time.Now().Unix())
	return int16(rand.Intn(int(max)-int(min)) + int(min))
}
