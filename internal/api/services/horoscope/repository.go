package horoscope

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
	findHoroscopes      *sqlx.Stmt
	findHoroscopeByName *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.HoroscopeRepository {
	stmts := Statement{
		findHoroscopes:      datasources.Prepare(dbReader, findHoroscopes),
		findHoroscopeByName: datasources.Prepare(dbReader, findHoroscopeByName),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindHoroscopes() (horoscopes []entities.Horoscope, err error) {
	err = r.stmt.findHoroscopes.Select(&horoscopes)
	if err != nil {
		log.Println("error while find horoscopes ", err)
	}

	return
}

func (r Repository) FindHoroscopeByName(horoscopeName string) (horoscope entities.Horoscope, err error) {
	err = r.stmt.findHoroscopeByName.Get(&horoscope, horoscopeName)
	if err != nil {
		log.Println("error while find horoscope by name ", err)
	}

	return
}
