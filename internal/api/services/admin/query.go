package admin

const (
	findAdminByEmail = `
		SELECT
			id_admin,
			fullname as full_name,
			email,
			password
		FROM 
			admin 
		WHERE 
			email = ? 
		LIMIT 1
	`

	refreshToken = `UPDATE admin SET token = ? WHERE email = ?`

	findAdminLoggedIn = `
		SELECT
			id_admin,
			fullname as full_name,
			email,
			password
		FROM 
			admin 
		WHERE 
			id_admin = ? AND token = ?
	`
)
