package models

type Game struct {
	Name              string    `json:"game_name"`
	Authors           string    `json:"authors"`
	Description       string    `json:"description"`
	Complexity        int       `json:"complexity"`
	PlayersPerSession int       `json:"players_per_session"`
	MastersPerSession int       `json:"masters_per_session"`
	Roles             []string  `json:"roles"`
	State             GameState `json:"game_state"`
}
