package bloodTypePoint

const (
	findBloodTypePoint = `
		SELECT
			b.id_blood_type as blood_type_id,
			c.id_category as category_id,
			b.name as blood_type_name,
			IFNULL(bp.point, "") as point
		FROM 
		    blood_type as b
		LEFT JOIN blood_type_point as bp
			ON bp.id_blood_type = b.id_blood_type
		LEFT JOIN category as c
			ON c.id_category = bp.id_category
		WHERE
		    c.id_category = ?
		ORDER BY b.name ASC
	`

	updateBloodTypePoint = `
		UPDATE blood_type_point SET point = :point WHERE id_blood_type = :id_blood_type AND id_category = :id_category 
	`
)
