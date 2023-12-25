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

	findAdminById = `
		SELECT
			id_admin as admin_id,
			fullname as full_name,
			email,
			IFNULL(phone_number, '') as phone_number,
			password
		FROM 
			admin 
		WHERE 
			id_admin = ? 
		LIMIT 1
	`

	findAdminByEmail = `
		SELECT
			id_admin,
			fullname as full_name,
			email,
			password,
			is_super_admin
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
			password,
			is_super_admin
		FROM 
			admin 
		WHERE 
			id_admin = ? AND token = ?
	`

	countEmail = `
		SELECT
			count(email) as total
		FROM
		    admin
		WHERE
		    email = ?
	`

	countPhoneNumber = `
		SELECT
			count(phone_number) as total
		FROM
		    admin
		WHERE
		    phone_number = ?
	`

	insertAdmin = `
		INSERT INTO admin 
		    (fullname, email, password, phone_number, is_super_admin)
		VALUES
		    (:fullname, :email, :password, :phone_number, :is_super_admin)
	`

	insertAccessMenu = `
		INSERT INTO access
			(id_menu, id_admin, status)
		VALUES
		    (:id_menu, :id_admin, :status)
	`

	updateAdmin = `
		UPDATE admin SET fullname = :fullname, email = :email, phone_number = :phone_number WHERE id_admin = :id_admin
	`

	deleteAccessMenu = `
		DELETE FROM access WHERE id_admin = :id_admin
	`

	countAdmin = `
		SELECT
			count(id_admin) as total
		FROM
		    admin
		WHERE
		    id_admin = ?
	`

	countAdminWithToken = `
		SELECT
			count(id_admin) as total
		FROM
		    admin
		WHERE
		    id_admin = ? AND
		    token = ?
	`

	deleteAdmin = `
		DELETE FROM admin WHERE id_admin = :id_admin
	`

	changePassword = `
		UPDATE
			admin
		SET
		    password = ?
		WHERE
		    id_admin = ?
	`
)
