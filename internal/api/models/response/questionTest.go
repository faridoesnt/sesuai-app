package response

import "Sesuai/internal/api/entities"

type QuestionTest struct {
	Questions []entities.QuestionTest `json:"question_list"`
	Answers   []entities.PointAnswer  `json:"answer"`
	Total     int                     `json:"total"`
}
