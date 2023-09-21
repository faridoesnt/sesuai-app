package entities

type Question struct {
	Id          string `db:"id_question"`
	CategoryId  string `db:"id_category"`
	Category    string `db:"category"`
	Photo       string `db:"photo"`
	QuestionIna string `db:"question_ina"`
	QuestionEn  string `db:"question_eng"`
}

type RequestQuestion struct {
	CategoryId  string `db:"id_category" json:"id_category"`
	QuestionIna string `db:"question_ina" json:"question_ina"`
	QuestionEn  string `db:"question_eng" json:"question_eng"`
	AdminId     string `db:"admin_id"`
}
