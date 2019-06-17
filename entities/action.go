package entities

import "fmt"

type ActionType uint

const (
	PickCharacter ActionType = 0
	PlayCharacter ActionType = 1
)

func ActionTypeAsString(a ActionType) string {
	switch a {
	case PickCharacter: return "pick-character"
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

type PickCharacterAction struct {
	GameResource GameResource
}

func (a PickCharacterAction) Resource() GameResource {
	return a.GameResource
}

func (a PickCharacterAction) ActionType() ActionType {
	return PickCharacter
}

func (a PickCharacterAction) AsString() string {
	return fmt.Sprintf("action:%s, %s",
		ActionTypeAsString(a.ActionType()),
		a.GameResource.AsString())
}

func NewPickCharacterAction(r GameResource) Action {
	return PickCharacterAction{
		GameResource: r,
	}
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


