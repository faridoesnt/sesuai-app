package result

const (
	findResult = `
		SELECT
		    s.id_submission as id,
			e.name as element_name,
			e.photo as element_image,
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
			s.id_submission as id,
			e.name as element_name,
			e.photo as element_image,
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
