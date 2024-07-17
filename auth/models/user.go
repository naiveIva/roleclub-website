package models

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	FatherName     string    `json:"father_name"`
	TelNumber      string    `json:"tel_number"`
	Password       string    `json:"password_hash"`
	IsHSEStudent   bool      `json:"is_hse_student"`
	IsBanned       bool      `json:"is_banned"`
	PlayedGames    int       `json:"played_games"`
	ConductedGames int       `json:"conducted_games"`
	Status         string    `json:"status"`
}
