package models

import "errors"

type GameState int

const (
	Playable GameState = iota
	Unplayable
	None
)

func (g *GameState) String() string {
	var str string
	switch *g {
	case Playable:
		str = "playable"
	case Unplayable:
		str = "unplayable"
	}
	return str
}

func StringToGameState(s string) (GameState, error) {
	switch s {
	case "playable":
		return Playable, nil
	case "unplayable":
		return Unplayable, nil
	}
	return None, errors.New("invalid game state")
}