package submission

const (
	findSubmissions = `
		SELECT
			s.id_submission as submission_id,
			u.fullname as name,
			u.email,
			s.time as timer,
			COUNT(ss.id_submission) as total_submission,
			COUNT(q.id_question) as total_question,
			s.created_at
		FROM 
			submission s
		LEFT JOIN user u
			ON s.id_user = u.id_user
		LEFT JOIN sub_submission ss
			ON s.id_submission = ss.id_submission
		LEFT JOIN question q
			ON q.id_question = ss.id_question
		ORDER BY s.created_at DESC
	`

	findResultSubmission = `
		SELECT
			c.id_category as category_id,
			c.name as category_name,
			ps.point as point
		FROM 
		    point_submission ps
		LEFT JOIN category c
			ON c.id_category = ps.id_category
		WHERE
		    ps.id_submission = ?
	`
)
