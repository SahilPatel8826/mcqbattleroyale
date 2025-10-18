package models

import "time"

type Room struct {
	RoomID    int       `json:"room_id" db:"room_id"`
	Code      string    `json:"code" db:"code"`
	Status    string    `json:"status" db:"status"` // waiting | active | ended
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
