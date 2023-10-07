package admin

const (
	findAdmins = `
		SELECT
			id_admin as admin_id,
			fullname as full_name,
			email,
			IFNULL(phone_number, '') as phone_number
		FROM
		    admin
		WHERE
		    is_super_admin != 1
		ORDER BY fullname ASC
	`

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
