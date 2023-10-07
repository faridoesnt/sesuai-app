package generateToken

const (
	findGenerateToken = `
				SELECT
				    g.id_token,
				    a.fullname as admin,
				    g.token,
				    g.status,
				    g.created_at
				FROM 
				    generate_token as g
				LEFT JOIN admin as a 
					ON a.id_admin = g.id_admin
				WHERE
				    g.id_admin = ?
				ORDER BY created_at DESC
	`

	insertNewToken = `
		INSERT INTO generate_token (
			id_admin, token, status
		) VALUES (
		    ?, ?, 'active'
		)
	`

	updateToken = `
		UPDATE
			generate_token
		SET
		    status = 'non active',
		    updated_at = CURRENT_TIMESTAMP
		WHERE
		    id_token = ?
	`
)
