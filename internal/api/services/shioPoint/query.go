package shioPoint

const (
	findShioPoint = `
		SELECT
			s.id_shio as shio_id,
			c.id_category as category_id,
			s.name as shio_name,
			IFNULL(sp.point, "") as point
		FROM 
		    shio as s
		LEFT JOIN shio_point as sp
			ON sp.id_shio = s.id_shio
		LEFT JOIN category as c
			ON c.id_category = sp.id_category
		WHERE
		    c.id_category = ?
		ORDER BY s.name ASC
	`
)
