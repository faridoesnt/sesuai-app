package pointAnswer

const (
	findPointAnswer = `
		SELECT
			id_answer as id,
			name,
			point
		FROM
		    answer
	`

	updatePointAnswer = `
		UPDATE answer SET point = :point WHERE id_answer = :id_answer 
	`

	findPointAnswerById = `
		SELECT
			id_answer as id,
			name,
			point
		FROM
		    answer
		WHERE
		    id_answer = ?
	`
)
