package accessMenu

const (
	findAccessMenuByAdminId = `
		SELECT
			IFNULL(m.menu, "") as menu_name
		FROM
		    access as a
		LEFT JOIN menu as m
			ON a.id_menu = m.id_menu
		WHERE
		    a.id_admin = ?
		ORDER BY m.menu ASC
	`

	countAdminAccessMenu = `
		SELECT
			count(a.id_admin) as count
		FROM
		    menu as m
		LEFT JOIN access as a
			ON a.id_menu = m.id_menu
		WHERE
		    a.id_admin = ? AND
		    m.menu = ?
		LIMIT 1
	`
)
