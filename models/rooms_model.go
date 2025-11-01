package models

import (
	"fmt"
	"time"

	middleware "mcqgame/db"
)

type Room struct {
	RoomID    int        `json:"room_id" db:"room_id"`
	Code      string     `json:"code" db:"code"`
	Status    string     `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	Questions []Question `json:"questions,omitempty"`
}

// ✅ Create a room with question stored as JSON
func CreateRoom(room Room) (Room, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	sqlStatement := `
		INSERT INTO rooms (code, status,created_at)
		VALUES ($1, $2, $3)
		RETURNING room_id
	`

	err := db.QueryRow(sqlStatement, room.Code, room.Status, time.Now()).
		Scan(&room.RoomID)
	if err != nil {
		return room, fmt.Errorf("unable to create room: %v", err)
	}

	sqlQuestion := `
		INSERT INTO questions (room_id, text, option_a, option_b, option_c, option_d, correct_option)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING question_id
	`

	for i := range room.Questions {
		q := &room.Questions[i]
		err = db.QueryRow(sqlQuestion, room.RoomID, q.Text, q.OptionA, q.OptionB, q.OptionC, q.OptionD, q.CorrectOption).
			Scan(&q.QuestionID)
		if err != nil {
			return room, fmt.Errorf("unable to insert question: %v", err)
		}
		q.RoomID = room.RoomID
	}

	return room, nil
}

// ✅ Fetch a room by ID (including question JSON)

func GetRoomWithQuestions(roomID int) (Room, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	var room Room
	room.RoomID = roomID

	// Step 1️⃣ Get room basic info
	sqlRoom := `SELECT code, status, created_at FROM rooms WHERE room_id=$1`
	err := db.QueryRow(sqlRoom, roomID).Scan(&room.Code, &room.Status, &room.CreatedAt)
	if err != nil {
		return room, fmt.Errorf("room not found: %v", err)
	}

	// Step 2️⃣ Get all related questions
	sqlQuestions := `
		SELECT question_id, room_id, text, option_a, option_b, option_c, option_d, correct_option
		FROM questions WHERE room_id=$1
	`
	rows, err := db.Query(sqlQuestions, roomID)
	if err != nil {
		return room, fmt.Errorf("error fetching questions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var q Question
		err = rows.Scan(&q.QuestionID, &q.RoomID, &q.Text, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectOption)
		if err != nil {
			return room, err
		}
		room.Questions = append(room.Questions, q)
	}

	return room, nil
}
func UpdateRoomStatus(roomID int, status string) error {
	db := middleware.CreateConnection()
	defer db.Close()

	_, err := db.Exec(`UPDATE rooms SET status=$1 WHERE room_id=$2`, status, roomID)
	if err != nil {
		return fmt.Errorf("failed to update room status: %v", err)
	}
	return nil
}
