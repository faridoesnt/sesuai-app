package result

const (
	findResultBySubmissionId = `
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
		    s.id_user = ? AND
		    s.id_submission = ? AND 
		    e.id_category IN (22, 20, 25, 17, 24)
		ORDER BY CASE WHEN e.id_category = 22 THEN 1
					  WHEN e.id_category = 20 THEN 2
					  WHEN e.id_category = 25 THEN 3
					  WHEN e.id_category = 17 THEN 4
					  WHEN e.id_category = 24 THEN 5 END
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
