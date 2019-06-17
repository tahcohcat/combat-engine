package entities

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

func Abs(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func randomFloat32(min, max float32) float32 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float32() * (max - min)
}

func randomInt16(min, max int16) int16 {
	rand.Seed(time.Now().UnixNano())
	i := Abs(int16(rand.Intn(int(max)-int(min)) + int(min)))
	return i
}

func Filter(vs []Player, f func(Player) bool) []Player {
	vsf := make([]Player, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func tribeToString(t CharacterTribe) string {
	switch t {
	case MechTribe: return "Mech"
	case DragonTribe: return "Dragon"
	default: return "Unknown"}
}

func getRandomCharacter() Character {
	tribe := CharacterTribe(randomInt16(0, NumTribes))
	switch tribe {
	case MechTribe: return NewGeneticMech()
	case DragonTribe: return NewGeneticDragon()
	default:
		panic(fmt.Sprintf("Unexpected tribe number %d", tribe))
	}
}

func getRandomRegionType() RegionType {
	return RegionType(randomInt16(0, NumRegions - 1))
}

func RoundPhaseAsString(phase RoundPhase) string {
	switch phase {
	case MarketPhase: return "market-phase"
	case CharPlacePhase: return "character-placement-phase"
	case FightPhase: return "fight-phase"
	case RoundEndPhase: return "end-of-round-phase"
	}
	return "unknown-phase"
}

func GenerateResource(resType GameResourceType, name string) GameResource {

	uuidv4, err := uuid.NewV4()
	if err != nil || uuidv4 == nil {
		panic("Could not generate uuid")
	}
	return GameResource{
		Name: name,
		Id:   *uuidv4,
		Type: resType,
	}
}