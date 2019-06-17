package entities

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
)

type GameResourceType uint

const (
	ResCharacter GameResourceType = 1
	ResRegion GameResourceType = 2
)

func GameResourceTypeAsString(g GameResourceType) string {
	switch g {
	case ResCharacter: return "resource-character"
	case ResRegion: return "resrouce-region"
	default: { panic("Implement me") }
	}
}

type GameResource struct {
	Name string
	Id  uuid.UUID
	Type GameResourceType
}

func (g GameResource) AsString() string {
	return fmt.Sprintf("%s [%s] %s", g.Name, g.Id.String(), GameResourceTypeAsString(g.Type))
}