package usedToken

const (
	insertUsedToken = `
		INSERT INTO used_token (
			id_token, id_user
		) VALUES (
		    ?, ?
		)
	`
)
