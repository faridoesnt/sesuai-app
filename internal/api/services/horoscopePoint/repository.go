package horoscopePoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"Sesuai/pkg/asql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findHoroscopePoint                 *sqlx.Stmt
	updateHoroscopePoint               *sqlx.NamedStmt
	findHoroscopePointByIdAndElementId *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.HoroscopePointRepository {
	stmts := Statement{
		findHoroscopePoint:                 datasources.Prepare(dbReader, findHoroscopePoint),
		updateHoroscopePoint:               datasources.PrepareNamed(dbWriter, updateHoroscopePoint),
		findHoroscopePointByIdAndElementId: datasources.Prepare(dbReader, findHoroscopePointByIdAndElementId),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindHoroscopePoint(elementId string) (horoscopePoint []entities.HoroscopePoint, err error) {
	err = r.stmt.findHoroscopePoint.Select(&horoscopePoint, elementId)
	if err != nil {
		log.Println("error while find horoscope point ", err)
	}

	return
}

func (r Repository) UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	for index, horoscopeId := range params.HoroscopeId {
		data := map[string]interface{}{
			"point":        params.Point[index],
			"id_horoscope": horoscopeId,
			"id_category":  params.ElementId,
		}

		_, err = tx.NamedStmt(r.stmt.updateHoroscopePoint).Exec(data)
		if err != nil {
			log.Println("error while update horoscope point ", err)
		}
	}

	return
}

func (r Repository) FindHoroscopePointByIdAndElementId(horoscopeId, elementId string) (horoscopePoint entities.HoroscopePoint, err error) {
	err = r.stmt.findHoroscopePointByIdAndElementId.Get(&horoscopePoint, horoscopeId, elementId)
	if err != nil {
		log.Println("error while find horoscope point by id and element id ", err)
	}

	return
}
