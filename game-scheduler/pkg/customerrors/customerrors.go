package customerrors

import "errors"

var (
	ErrGameNotFound      = errors.New("game not found")
	ErrGameAlreadyExists = errors.New("game already exists")
	ErrEventNotFound     = errors.New("event not found")

	// ErrInvalidGameState  = errors.New("invalid game state")
)
