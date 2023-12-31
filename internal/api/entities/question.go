package entities

type Question struct {
	Id          string `db:"id_question"`
	ElementId   string `db:"element_id"`
	Element     string `db:"element"`
	Photo       string `db:"photo"`
	QuestionIna string `db:"question_ina"`
	QuestionEn  string `db:"question_eng"`
}

type RequestQuestion struct {
	ElementId   string `db:"element_id" json:"element_id"`
	QuestionIna string `db:"question_ina" json:"question_ina"`
	QuestionEn  string `db:"question_eng" json:"question_eng"`
	AdminId     string `db:"admin_id"`
}
