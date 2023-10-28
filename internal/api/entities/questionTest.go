package entities

type QuestionTest struct {
	QuestionId  string `db:"question_id" json:"question_id"`
	ElementId   string `db:"element_id" json:"element_id"`
	QuestionIna string `db:"question_ina" json:"question_ina"`
	QuestionEn  string `db:"question_eng" json:"question_eng"`
}
