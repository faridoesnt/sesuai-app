package result

const (
	findResult = `
		SELECT
			e.name as element_name,
			ps.point
		FROM
		    submission as s
		LEFT JOIN point_submission as ps 
			ON ps.id_submission = s.id_submission
		LEFT JOIN category as e
			ON e.id_category = ps.id_category
		WHERE
		    s.id_user = ?
		LIMIT 3
	`

	findAllResult = `
		SELECT
			e.name as element_name,
			ps.point
		FROM
		    submission as s
		LEFT JOIN point_submission as ps 
			ON ps.id_submission = s.id_submission
		LEFT JOIN category as e
			ON e.id_category = ps.id_category
		WHERE
		    s.id_user = ?
	`
)
