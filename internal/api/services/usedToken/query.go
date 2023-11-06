package usedToken

const (
	insertUsedToken = `
		INSERT INTO used_token (
			id_token, id_user
		) VALUES (
		    ?, ?
		)
	`

	countUserToken = `
		SELECT
			COUNT(ut.id_token) as count
		FROM
		    used_token as ut
		LEFT JOIN 
			generate_token as gt ON ut.id_token = gt.id_token
		WHERE
		    gt.token = ? AND
		    ut.id_user = ?
	`
)
