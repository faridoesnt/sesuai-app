package questionTest

const (
	findQuestionsTest = `
		SELECT
			id_question as question_id,
			id_category as element_id,
			question_ina,
			question_eng
		FROM
			question
		ORDER BY RAND()
	`
)
