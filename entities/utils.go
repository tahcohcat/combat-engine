package entities

import (
	"fmt"
	"math/rand"
	"time"
)

func randomFloat32(min, max float32) float32 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float32() * (max - min)
}

func randomInt16(min, max int16) int16 {
	rand.Seed(time.Now().UnixNano())
	return int16(rand.Intn(int(max)-int(min)) + int(min))
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
	tribe := CharacterTribe(randomInt16(0, NumTribes - 1))
	switch tribe {
	case MechTribe: return NewGeneticMech()
	case DragonTribe: return NewGeneticDragon()
	default:
		panic(fmt.Sprintf("Unexpected tribe number %d", tribe))
	}
}