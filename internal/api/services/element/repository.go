package element

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
	findElements         *sqlx.Stmt
	findElementById      *sqlx.Stmt
	insertElement        *sqlx.NamedStmt
	updateElement        *sqlx.Stmt
	deleteElement        *sqlx.Stmt
	insertShioPoint      *sqlx.NamedStmt
	insertHoroscopePoint *sqlx.NamedStmt
	insertBloodTypePoint *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ElementRepository {
	stmts := Statement{
		findElements:         datasources.Prepare(dbReader, findElements),
		findElementById:      datasources.Prepare(dbReader, findElementById),
		insertElement:        datasources.PrepareNamed(dbWriter, insertElement),
		updateElement:        datasources.Prepare(dbWriter, updateElement),
		deleteElement:        datasources.Prepare(dbWriter, deleteElement),
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

func (r Repository) FindElements() (categories []entities.Element, err error) {
	err = r.stmt.findElements.Select(&categories)
	if err != nil {
		log.Println("error while find element ", err)
	}

	return
}

func (r Repository) FindElementById(elementId string) (element entities.Element, err error) {
	err = r.stmt.findElementById.Get(&element, elementId)
	if err != nil {
		log.Println("error while find element by id ", err)
	}

	return
}

func (r Repository) InsertElement(element entities.RequestElement) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	res, err := tx.NamedStmt(r.stmt.insertElement).Exec(element)
	if err != nil {
		log.Println("error while insert element ", err)
	}

	resId, _ := res.LastInsertId()

	elementId := strconv.FormatInt(resId, 10)

	shios, err := r.service.Shio.GetShio()
	if len(shios) > 0 {
		for _, shio := range shios {
			data := map[string]interface{}{
				"id_shio":     shio.Id,
				"id_category": elementId,
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
				"id_category":  elementId,
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
				"id_category":   elementId,
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

func (r Repository) UpdateElement(elementId string, params entities.RequestElement) (err error) {
	_, err = r.stmt.updateElement.Exec(params.Name, params.FileName, params.AdminId, elementId)
	if err != nil {
		log.Println("error while update element ", err)
	}

	return
}

func (r Repository) DeleteElement(elementId string) (err error) {
	_, err = r.stmt.deleteElement.Exec(elementId)
	if err != nil {
		log.Println("error while delete element ", err)
	}

	return
}
