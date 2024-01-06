package usedToken

const (
	insertUsedToken = `
		INSERT INTO used_token (
			id_token, id_submission
		) VALUES (
		    ?, ?
		)
	`

	countSubmissionToken = `
		SELECT
			COUNT(ut.id_token) as count
		FROM
		    used_token as ut
		LEFT JOIN 
			generate_token as gt ON ut.id_token = gt.id_token
		WHERE
		    gt.token = ? AND
		    ut.id_submission = ?
	`

	findUsedTokenByUserId = `
		SELECT
			gt.token as token
		FROM
		    used_token as ut
		LEFT JOIN 
			generate_token as gt ON ut.id_token = gt.id_token
		WHERE
		    ut.id_user = ?
	`
)
