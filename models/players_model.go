package models

import "time"

type Player struct {
	player_id   int       `json:"player_id" db:"player_id"`
	name        string    `json:"name" db:"name"`
	email       string    `json:"email" db:"email"`
	total_score int       `json:"total_score`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
