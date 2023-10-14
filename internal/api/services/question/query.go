package question

const (
	findQuestionsByElementId = `
		SELECT
			q.id_question,
			c.id_category as element_id,
			c.name as element,
			c.photo,
			q.question_ina,
			q.question_eng
		FROM
		    question q
		LEFT JOIN category c
			ON q.id_category = c.id_category
		WHERE
		    c.id_category = ?
		ORDER BY q.created_at DESC
		LIMIT 3
	`

	findAllQuestionsByElementId = `
		SELECT
			q.id_question,
			c.id_category as element_id,
			c.name as element,
			c.photo,
			q.question_ina,
			q.question_eng
		FROM
		    question q
		LEFT JOIN category c
			ON q.id_category = c.id_category
		WHERE
		    c.id_category = ?
		ORDER BY q.created_at DESC
	`

	findQuestion = `
		SELECT
			q.id_question,
			c.id_category as element_id,
			c.name as element,
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
		    :element_id, :question_ina, :question_eng, :admin_id
		)
	`

	updateQuestion = `
		UPDATE question SET id_category = :element_id, question_ina = :question_ina, question_eng = :question_eng WHERE id_question = :id_question
	`

	deleteQuestion = `
		DELETE FROM question WHERE id_question = ?
	`
)
