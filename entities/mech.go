package entities

type Mech struct {
	base CharacterBase
}

func (m Mech) Display() string {
	return m.base.Display()
}

func (m Mech) Name() string {
	return m.base.Name
}

func (m Mech) Stats() Stats {
	return m.base.Stats
}

func (Mech) Tribe() CharacterTribe {
	return MechTribe
}

func (m Mech) Resource() GameResource {
	return m.base.Resource
}

func NewMech() Character {
	c  := Mech{}
	c.base.Name = "MechName"
	c.base.Tribe = MechTribe
	c.base.Stats = mechBaseStats()
	c.base.Resource = GenerateResource(ResCharacter, "resource-character-mech")
	return c
}

func mechBaseStats() Stats {
	return Stats {
		Health : 100,
		AttackDamage : 50,
		AttackSpeed : 1.0,
		MovementSpeed : 40,
		CritPercentage : 0.0,
	}
}

func NewGeneticMech() Character {

	character := NewMech()

	mech, ok := character.(Mech)
	if !ok {
		panic("We expected a mech")
	}

	mech.base.Stats.Health += uint16(randomInt16(-10, 10))
	mech.base.Stats.AttackDamage += uint16(randomInt16(-10, 10))
	mech.base.Stats.AttackSpeed += randomFloat32(-0.3, 0.3)
	mech.base.Stats.MovementSpeed += uint16(randomInt16(-10, 10))
	return &mech
}
