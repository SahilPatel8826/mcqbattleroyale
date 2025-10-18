package models

import "time"

type RoomPlayer struct {
	RoomPlayerID  int       `json:"room_player_id" db:"room_player_id"`
	RoomID        int       `json:"room_id" db:"room_id"`
	PlayerID      int       `json:"player_id" db:"player_id"`
	score         int       `json:"score" db:"score"`
	Is_Eliminated bool      `json:"is_eliminated" db:"is_eliminated"`
	Joined_At     time.Time `json:"joined_at" db:"joined_at"`
}
