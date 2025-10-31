package models

import (
	"fmt"
	middleware "mcqgame/db"
	"time"
)

type Answer struct {
	AnswerID       int       `json:"answer_id" db:"answer_id"`
	RoomID         int       `json:"room_id" db:"room_id"`
	QuestionID     int       `json:"question_id" db:"question_id"`
	PlayerID       int       `json:"player_id" db:"player_id"`
	SelectedOption string    `json:"selected_option" db:"selected_option"`
	IsCorrect      bool      `json:"is_correct" db:"is_correct"`
	SubmittedAt    time.Time `json:"submitted_at" db:"submitted_at"`
}

func SubmitAnswer(answer Answer) (Answer, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	// Step 1: Find correct option from the question
	var correctOption string
	err := db.QueryRow(`
		SELECT correct_option 
		FROM questions 
		WHERE question_id=$1
	`, answer.QuestionID).Scan(&correctOption)
	if err != nil {
		return answer, fmt.Errorf("invalid question or missing correct answer")
	}

	// Step 2: Check correctness
	if answer.SelectedOption == correctOption {
		answer.IsCorrect = true
	} else {
		answer.IsCorrect = false
	}

	// Step 3: Store answer in answers table
	err = db.QueryRow(`
		INSERT INTO answers (room_id, question_id, player_id, selected_option, is_correct) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING answer_id, submitted_at
	`, answer.RoomID, answer.QuestionID, answer.PlayerID, answer.SelectedOption, answer.IsCorrect).
		Scan(&answer.AnswerID, &answer.SubmittedAt)

	if err != nil {
		return answer, fmt.Errorf("failed to save answer: %v", err)
	}

	// Step 4: If correct → Increase score by 10
	if answer.IsCorrect {
		_, err := db.Exec(`
			UPDATE room_players 
			SET total_score = total_score + 10 
			WHERE player_id=$1
		`, answer.PlayerID)

		if err != nil {
			return answer, fmt.Errorf("failed to update score: %v", err)
		}
	}

	return answer, nil
}

func GetAnswersByPlayerID(playerID int, roomID int) ([]Answer, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	rows, err := db.Query(`SELECT answer_id, room_id, question_id, player_id, selected_option, is_correct, submitted_at 
	                       FROM answers WHERE player_id = $1 AND room_id=$2`, playerID, roomID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []Answer

	for rows.Next() {
		var a Answer
		rows.Scan(&a.AnswerID, &a.RoomID, &a.QuestionID, &a.PlayerID, &a.SelectedOption, &a.IsCorrect, &a.SubmittedAt)
		answers = append(answers, a)
	}
	return answers, nil
}
