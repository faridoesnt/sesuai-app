package category

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/services/bloodType"
	"Sesuai/internal/api/services/horoscope"
	"Sesuai/internal/api/services/shio"
	"Sesuai/pkg/asql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
	service  constracts.Services
}

type Statement struct {
	findCategory         *sqlx.Stmt
	findCategoryById     *sqlx.Stmt
	insertCategory       *sqlx.NamedStmt
	updateCategory       *sqlx.Stmt
	deleteCategory       *sqlx.Stmt
	insertShioPoint      *sqlx.NamedStmt
	insertHoroscopePoint *sqlx.NamedStmt
	insertBloodTypePoint *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.CategoryRepository {
	stmts := Statement{
		findCategory:         datasources.Prepare(dbReader, findCategory),
		findCategoryById:     datasources.Prepare(dbReader, findCategoryById),
		insertCategory:       datasources.PrepareNamed(dbWriter, insertCategory),
		updateCategory:       datasources.Prepare(dbWriter, updateCategory),
		deleteCategory:       datasources.Prepare(dbWriter, deleteCategory),
		insertShioPoint:      datasources.PrepareNamed(dbWriter, insertShioPoint),
		insertHoroscopePoint: datasources.PrepareNamed(dbWriter, insertHoroscopePoint),
		insertBloodTypePoint: datasources.PrepareNamed(dbWriter, insertBloodTypePoint),
	}

	app := constracts.App{
		Datasources: &constracts.Datasources{
			WriterDB: dbWriter,
			ReaderDB: dbReader,
		},
	}

	svc := constracts.Services{
		Shio:      shio.Init(&app),
		Horoscope: horoscope.Init(&app),
		BloodType: bloodType.Init(&app),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
		service:  svc,
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
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	res, err := tx.NamedStmt(r.stmt.insertCategory).Exec(category)
	if err != nil {
		log.Println("error while insert category ", err)
	}

	resId, _ := res.LastInsertId()

	categoryId := strconv.FormatInt(resId, 10)

	shios, err := r.service.Shio.GetShio()
	if len(shios) > 0 {
		for _, shio := range shios {
			data := map[string]interface{}{
				"id_shio":     shio.Id,
				"id_category": categoryId,
				"point":       "0",
			}

			_, err = tx.NamedStmt(r.stmt.insertShioPoint).Exec(data)
			if err != nil {
				log.Println("error while insert shio point")
			}
		}
	}

	horoscopes, err := r.service.Horoscope.GetHoroscopes()
	if len(horoscopes) > 0 {
		for _, horoscope := range horoscopes {
			data := map[string]interface{}{
				"id_horoscope": horoscope.Id,
				"id_category":  categoryId,
				"point":        "0",
			}

			_, err = tx.NamedStmt(r.stmt.insertHoroscopePoint).Exec(data)
			if err != nil {
				log.Println("error while insert horoscope point ", err)
			}
		}
	}

	bloodTypes := r.service.BloodType.GetBloodType()
	if len(bloodTypes) > 0 {
		for _, bloodType := range bloodTypes {
			data := map[string]interface{}{
				"id_blood_type": bloodType.Id,
				"id_category":   categoryId,
				"point":         "0",
			}

			_, err = tx.NamedStmt(r.stmt.insertBloodTypePoint).Exec(data)
			if err != nil {
				log.Println("error while blood type point ", err)
			}
		}
	}

	return
}

func (r Repository) UpdateCategory(categoryId string, params entities.RequestCategory) (err error) {
	_, err = r.stmt.updateCategory.Exec(params.Name, params.FileName, params.AdminId, categoryId)
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
