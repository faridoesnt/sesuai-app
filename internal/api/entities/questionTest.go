package entities

type QuestionTest struct {
	QuestionId  string `db:"question_id" json:"question_id"`
	ElementId   string `db:"element_id" json:"element_id"`
	QuestionIna string `db:"question_ina" json:"question_ina"`
	QuestionEn  string `db:"question_eng" json:"question_eng"`
}

type SubmitQuestionTestItem struct {
	QuestionId string `json:"question_id"`
	ElementId  string `json:"element_id"`
	AnswerId   string `json:"answer_id"`
}

type SubmitQuestionTest struct {
	Submit []SubmitQuestionTestItem `json:"submit"`
	Timer  string                   `json:"timer"`
}
