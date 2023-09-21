package category

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findCategory     *sqlx.Stmt
	findCategoryById *sqlx.Stmt
	insertCategory   *sqlx.NamedStmt
	updateCategory   *sqlx.Stmt
	deleteCategory   *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.CategoryRepository {
	stmts := Statement{
		findCategory:     datasources.Prepare(dbReader, findCategory),
		findCategoryById: datasources.Prepare(dbReader, findCategoryById),
		insertCategory:   datasources.PrepareNamed(dbWriter, insertCategory),
		updateCategory:   datasources.Prepare(dbWriter, updateCategory),
		deleteCategory:   datasources.Prepare(dbWriter, deleteCategory),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindCategory() (categories []entities.Category, err error) {
	err = r.stmt.findCategory.Select(&categories)
	if err != nil {
		log.Println("error while find category ", err)
	}

	return
}

func (r Repository) FindCategoryById(categoryId string) (category entities.Category, err error) {
	err = r.stmt.findCategoryById.Get(&category, categoryId)
	if err != nil {
		log.Println("error while find category by id ", err)
	}

	return
}

func (r Repository) InsertCategory(category entities.RequestCategory) (err error) {
	_, err = r.stmt.insertCategory.Exec(category)
	if err != nil {
		log.Println("error while insert category ", err)
	}

	return
}

func (r Repository) UpdateCategory(categoryId string, params entities.RequestCategory) (err error) {
	_, err = r.stmt.updateCategory.Exec(params.Name, params.Photo, params.AdminId, categoryId)
	if err != nil {
		log.Println("error while update category ", err)
	}

	return
}

func (r Repository) DeleteCategory(categoryId string) (err error) {
	_, err = r.stmt.deleteCategory.Exec(categoryId)
	if err != nil {
		log.Println("error while delete category ", err)
	}

	return
}
