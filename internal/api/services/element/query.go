package element

const (
	findElements = `
				SELECT
				    id_category,
				    name,
				    IFNULL(photo, "") as photo
				FROM 
				    category
				ORDER BY name ASC
	`

	findElementById = `
				SELECT
				    id_category,
				    name,
				    IFNULL(photo, "") as photo
				FROM 
				    category
				WHERE
				    id_category = ?
	`

	insertElement = `
		INSERT INTO category (
		    name, photo, created_by
		) VALUES (
		    :name, :photo, :admin_id
		)
	`

	updateElement = `
		UPDATE
			category
		SET
		    name = ?, photo = ?, updated_by = ?, updated_at = current_timestamp
		WHERE
		    id_category = ?
	`

	deleteElement = `
		DELETE FROM category WHERE id_category = ?
	`

	insertShioPoint = `
		INSERT INTO shio_point (id_shio, id_category, point) VALUES (:id_shio, :id_category, :point)
	`

	insertHoroscopePoint = `
		INSERT INTO horoscope_point (id_horoscope, id_category, point) VALUES (:id_horoscope, :id_category, :point)
	`

	insertBloodTypePoint = `
		INSERT INTO blood_type_point (id_blood_type, id_category, point) VALUES (:id_blood_type, :id_category, :point)
	`
)
