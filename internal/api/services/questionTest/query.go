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

	insertSubmission = `
		INSERT INTO submission (id_user, time) VALUES (:id_user, :time)
	`

	insertSubSubmission = `
		INSERT INTO sub_submission (id_submission, id_question, id_answer) VALUES (:id_submission, :id_question, :id_answer)
	`

	insertPointSubmission = `
		INSERT INTO point_submission (id_submission, id_category, point) VALUES (:id_submission, :id_category, :point)
	`
)
