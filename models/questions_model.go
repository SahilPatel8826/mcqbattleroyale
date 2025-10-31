package models

import (
	"fmt"
	middleware "mcqgame/db"
)

type Question struct {
	QuestionID    int    `json:"question_id" db:"question_id"`
	RoomID        int    `json:"room_id" db:"room_id"`
	Text          string `json:"text" db:"text"`
	OptionA       string `json:"option_a" db:"option_a"`
	OptionB       string `json:"option_b" db:"option_b"`
	OptionC       string `json:"option_c" db:"option_c"`
	OptionD       string `json:"option_d" db:"option_d"`
	CorrectOption string `json:"correct_answer" db:"correct_option"`
}

func CreateQuestion(question Question) (Question, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO questions (room_id,text,option_a,option_b,option_c,option_d,correct_option) values ($1,$2,$3,$4,$5,$6,$7)   RETURNING question_id;`
	err := db.QueryRow(sqlStatement, question.RoomID, question.Text, question.OptionA, question.OptionB, question.OptionC, question.OptionD, question.CorrectOption).Scan(&question.QuestionID)
	if err != nil {
		fmt.Println("error in insert in table", err)
	}
	return question, nil
}
func GetQuestionsByRoomID(roomID int) ([]Question, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	rows, err := db.Query(`SELECT question_id, room_id, text, option_a, option_b, option_c, option_d, correct_option 
	                       FROM questions WHERE room_id = $1`, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question

	for rows.Next() {
		var q Question
		rows.Scan(&q.QuestionID, &q.RoomID, &q.Text, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectOption)
		questions = append(questions, q)
	}

	return questions, nil
}
func UpdateQuestionByID(id int, q Question) (Question, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	// Step 1: Fetch current question
	var current Question
	err := db.QueryRow(`
		SELECT question_id, room_id, text, option_a, option_b, option_c, option_d, correct_option
		FROM questions
		WHERE question_id = $1
	`, id).Scan(
		&current.QuestionID,
		&current.RoomID,
		&current.Text,
		&current.OptionA,
		&current.OptionB,
		&current.OptionC,
		&current.OptionD,
		&current.CorrectOption,
	)
	if err != nil {
		return Question{}, fmt.Errorf("question not found: %v", err)
	}

	// Step 2: Only update non-empty fields (Go strings can't be nil, so we check != "")
	if q.RoomID != 0 {
		current.RoomID = q.RoomID
	}
	if q.Text != "" {
		current.Text = q.Text
	}
	if q.OptionA != "" {
		current.OptionA = q.OptionA
	}
	if q.OptionB != "" {
		current.OptionB = q.OptionB
	}
	if q.OptionC != "" {
		current.OptionC = q.OptionC
	}
	if q.OptionD != "" {
		current.OptionD = q.OptionD
	}
	if q.CorrectOption != "" {
		current.CorrectOption = q.CorrectOption
	}

	// Step 3: Update in DB
	sqlStatement := `
		UPDATE questions
		SET room_id = $1, text = $2, option_a = $3, option_b = $4, option_c = $5, option_d = $6, correct_option = $7
		WHERE question_id = $8
		RETURNING question_id, room_id, text, option_a, option_b, option_c, option_d, correct_option;
	`

	err = db.QueryRow(
		sqlStatement,
		current.RoomID,
		current.Text,
		current.OptionA,
		current.OptionB,
		current.OptionC,
		current.OptionD,
		current.CorrectOption,
		id,
	).Scan(
		&current.QuestionID,
		&current.RoomID,
		&current.Text,
		&current.OptionA,
		&current.OptionB,
		&current.OptionC,
		&current.OptionD,
		&current.CorrectOption,
	)
	if err != nil {
		return Question{}, fmt.Errorf("update failed: %v", err)
	}

	return current, nil
}
func DeleteQuestionByID(id int) error {
	db := middleware.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM questions WHERE question_id=$1`
	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		return fmt.Errorf("failed to delete question: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no question found with ID %d", id)
	}

	fmt.Println("✅ Question deleted with ID:", id)
	return nil
}
