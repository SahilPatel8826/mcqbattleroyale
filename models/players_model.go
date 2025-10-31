package models

import (
	"fmt"
	middleware "mcqgame/db"
	"time"
)

type Player struct {
	PlayerID   int       `json:"player_id" db:"player_id"`
	RoomID     int       `json:"room_id" db:"room_id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	TotalScore int       `json:"total_score" db:"total_score"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func JoinRoom(roomCode string, player Player) (Player, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	// Step 1: Check if room exists and is joinable
	var roomID int
	err := db.QueryRow(`SELECT room_id FROM rooms WHERE code=$1 AND status='waiting'`, roomCode).Scan(&roomID)
	if err != nil {
		return player, fmt.Errorf("room not found or not joinable")
	}

	// Step 2: Insert player into room
	sql := `
		INSERT INTO room_players (room_id, name, email, total_score)
		VALUES ($1, $2, $3, 0)
		RETURNING player_id, created_at
	`
	err = db.QueryRow(sql, roomID, player.Name, player.Email).Scan(&player.PlayerID, &player.CreatedAt)
	if err != nil {
		return player, fmt.Errorf("failed to join room: %v", err)
	}

	player.RoomID = roomID
	player.TotalScore = 0

	return player, nil
}

func GetPlayersByRoomID(roomID int) ([]Player, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	query := `SELECT player_id, room_id, name, email, total_score, created_at 
              FROM room_players 
              WHERE room_id = $1`

	rows, err := db.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player

	for rows.Next() {
		var p Player
		if err := rows.Scan(
			&p.PlayerID,
			&p.RoomID,
			&p.Name,
			&p.Email,
			&p.TotalScore,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}
func DeletePlayerByID(playerID int) (bool, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM room_players WHERE player_id = $1`
	rows, err := db.Exec(sqlStatement, playerID)
	if err != nil {
		return false, err

	}
	rowsAffected, _ := rows.RowsAffected()
	return rowsAffected > 0, nil

}
