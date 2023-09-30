package horoscopePoint

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
	findHoroscopePoint *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.HoroscopePointRepository {
	stmts := Statement{
		findHoroscopePoint: datasources.Prepare(dbReader, findHoroscopePoint),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindHoroscopePoint(categoryId string) (horoscopePoint []entities.HoroscopePoint, err error) {
	err = r.stmt.findHoroscopePoint.Select(&horoscopePoint, categoryId)
	if err != nil {
		log.Println("error while find horoscope point ", err)
	}

	return
}
