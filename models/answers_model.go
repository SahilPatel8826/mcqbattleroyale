package models

import "time"

type Answer struct {
	AnswerID    int       `json:"answer_id" db:"answer_id"`
	RoomID      int       `json:"room_id" db:"room_id"`
	QuestionID  int       `json:"question_id" db:"question_id"`
	PlayerID    int       `json:"player_id" db:"player_id"`
	AnswerText  string    `json:"submited_answer" db:"answer_text"`
	IsCorrect   bool      `json:"is_correct" db:"is_correct"`
	SubmittedAt time.Time `json:"submitted_at" db:"submitted_at"`
}
