package menu

const (
	findMenuIdByName = `
		SELECT
			id_menu as id,
			menu
		FROM
		    menu
		WHERE
		    menu = ?
	`
)
