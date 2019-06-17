package entities

import "fmt"


type Market struct {
	CharacterOptions []Character
	Size uint8
}

func (m *Market) Populate() {

	m.CharacterOptions = nil

	for i := uint8(0); i <= m.Size; i++ {
		m.CharacterOptions = append(
			m.CharacterOptions,
			getRandomCharacter())
	}
}

func (m Market) Display() string {

	s := string("")
	for _, character := range m.CharacterOptions {
		s += fmt.Sprintf("%s", character.Display())
	}
	return fmt.Sprintf("%s\n", s)
}

