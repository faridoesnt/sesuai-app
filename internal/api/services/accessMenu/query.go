package accessMenu

const (
	findAccessMenuByAdminId = `
		SELECT
			a.id_access as access_id,
			a.id_menu as menu_id,
			m.menu as menu_name,
			a.status
		FROM
		    access as a
		LEFT JOIN menu as m
			ON a.id_menu = m.id_menu
		WHERE
		    a.id_admin = ?
		ORDER BY m.menu ASC
	`
)
