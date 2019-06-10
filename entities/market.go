package entities

import "fmt"


type Market struct {
	characterOptions []Character
	size uint8
}

func (m *Market) Populate() {

	for i := uint8(0); i <= m.size; i++ {
		m.characterOptions = append(
			m.characterOptions,
			getRandomCharacter())
	}
}

func (m Market) Display() string {

	s := string("")
	for _, character := range m.characterOptions {
		s += fmt.Sprintf("%s", character.Display())
	}
	return fmt.Sprintf("%s\n", s)
}

