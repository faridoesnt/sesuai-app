package bloodType

const (
	findBloodType = `
				SELECT
				    id_blood_type,
				    name
				FROM 
				    blood_type
	`

	findBloodTypeUser = `
		SELECT
			bt.id_blood_type,
			bt.name
		FROM
		    blood_type as bt
		LEFT JOIN user as u
			ON bt.id_blood_type = u.id_blood_type
		WHERE
		    u.id_user = ?
	`

	countBloodTypeById = `
		SELECT
			COUNT(id_blood_type) as count
		FROM
		    blood_type
		WHERE
		    id_blood_type = ?
	`
)
