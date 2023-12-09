package shio

const (
	findShio = `
		SELECT
			id_shio,
			name
		FROM
		    shio
		ORDER BY id_shio ASC
	`

	findShioUser = `
		SELECT
			s.id_shio,
			s.name
		FROM
		    shio as s
		LEFT JOIN user as u
			ON s.id_shio = u.id_shio
		WHERE
		    u.id_user = ?
	`

	countShioById = `
		SELECT
			COUNT(id_shio) as count
		FROM
		    shio
		WHERE
		    id_shio = ?
	`
)
