package category

const (
	findCategory = `
				SELECT
				    id_category,
				    name,
				    IFNULL(photo, "") as photo
				FROM 
				    category
	`

	findCategoryById = `
				SELECT
				    id_category,
				    name,
				    IFNULL(photo, "") as photo
				FROM 
				    category
				WHERE
				    id_category = ?
	`

	insertCategory = `
		INSERT INTO category (
		    name, photo, created_by
		) VALUES (
		    :name, :photo, :admin_id
		)
	`

	updateCategory = `
		UPDATE
			category
		SET
		    name = ?, photo = ?, updated_by = ?, updated_at = current_timestamp
		WHERE
		    id_category = ?
	`

	deleteCategory = `
		DELETE FROM category WHERE id_category = ?
	`
)
