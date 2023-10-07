package question

const (
	findQuestionsByCategoryId = `
		SELECT
			q.id_question,
			c.id_category,
			c.name as category,
			c.photo,
			q.question_ina,
			q.question_eng
		FROM
		    question q
		LEFT JOIN category c
			ON q.id_category = c.id_category
		WHERE
		    c.id_category = ?
	`

	findQuestion = `
		SELECT
			q.id_question,
			c.id_category,
			c.name as category,
			c.photo,
			q.question_ina,
			q.question_eng
		FROM
		    question q
		LEFT JOIN category c
			ON q.id_category = c.id_category
		WHERE q.id_question = ?
	`

	insertQuestion = `
		INSERT INTO question (
		    id_category, question_ina, question_eng, created_by
		) VALUES (
		    :id_category, :question_ina, :question_eng, :admin_id
		)
	`

	deleteQuestion = `
		DELETE FROM question WHERE id_question = ?
	`
)
