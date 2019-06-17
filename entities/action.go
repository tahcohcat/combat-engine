package entities

import "fmt"

type ActionType uint

const (
	PlayCharacter ActionType = 0
)

func ActionTypeAsString(a ActionType) string {
	switch a {
	case PlayCharacter: return "play-character"
	default: {
		panic("Implement me")
	}
}
}

type Action interface {
	Resource() GameResource
	ActionType() ActionType
	AsString() string
}

type PlayCharacterAction struct {
	GameResource GameResource
}

func (a PlayCharacterAction) Resource() GameResource {
	return a.GameResource
}

func (a PlayCharacterAction) ActionType() ActionType {
	return PlayCharacter
}

func (a PlayCharacterAction) AsString() string {
	return fmt.Sprintf("action:%s, %s",
		ActionTypeAsString(a.ActionType()),
		a.GameResource.AsString())
}

func NewPlayCharacterAction(r GameResource) Action {
	return PlayCharacterAction{
		GameResource: r,
	}
}
