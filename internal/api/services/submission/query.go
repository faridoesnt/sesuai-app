package submission

const (
	findSubmissions = `
		SELECT
			sub.id_submission as submission_id,
			u.fullname as name,
			u.email,
			s.time as timer,
			sub.total_submission,
			sub.total_question,
			s.created_at,
			IFNULL(gt.token, "-") as token
		FROM (
		    SELECT
		        s.id_submission,
			    COUNT(ss.id_submission) as total_submission,
			    COUNT(q.id_question) as total_question
		    FROM
			    submission s
		    LEFT JOIN sub_submission ss
		    	ON s.id_submission = ss.id_submission
		    LEFT JOIN question q
		    	ON q.id_question = ss.id_question
            GROUP BY s.id_submission
        ) sub
		LEFT JOIN submission s
		    ON sub.id_submission = s.id_submission
		LEFT JOIN user u
		    ON s.id_user = u.id_user
		LEFT JOIN used_token ut
			ON u.id_user = ut.id_user
		LEFT JOIN generate_token as gt
		    ON ut.id_token = gt.id_token
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
