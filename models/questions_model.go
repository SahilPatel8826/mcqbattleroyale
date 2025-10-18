package models

type Question struct {
	QuestionID    int      `json:"question_id" db:"question_id"`
	Text          string   `json:"text" db:"text"`
	Options       []string `json:"options" db:"options"`
	OptionA       string   `json:"option_a" db:"option_a"`
	OptionB       string   `json:"option_b" db:"option_b"`
	OptionC       string   `json:"option_c" db:"option_c"`
	OptionD       string   `json:"option_d" db:"option_d"`
	CorrectOption string   `json:"correct_answer" db:"correct_option"`
}
