package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                 uuid.UUID `json:"id"`
	Gamename           string    `json:"gamename"`
	Date               time.Time `json:"date"`
	NumOfSessions      int       `json:"num_of_sessions"`
	IsSubscriptionOpen bool      `json:"is_subscription_open"`
}
