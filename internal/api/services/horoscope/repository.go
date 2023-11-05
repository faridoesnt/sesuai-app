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
	findHoroscopeUser   *sqlx.Stmt
	countHoroscopeById  *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.HoroscopeRepository {
	stmts := Statement{
		findHoroscopes:      datasources.Prepare(dbReader, findHoroscopes),
		findHoroscopeByName: datasources.Prepare(dbReader, findHoroscopeByName),
		findHoroscopeUser:   datasources.Prepare(dbReader, findHoroscopeUser),
		countHoroscopeById:  datasources.Prepare(dbReader, countHoroscopeById),
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

func (r Repository) FindHoroscopeUser(userId string) (horoscope entities.Horoscope, err error) {
	err = r.stmt.findHoroscopeUser.Get(&horoscope, userId)
	if err != nil {
		log.Println("error while find horoscope user ", err)
	}

	return
}

func (r Repository) CountHoroscopeById(horoscopeId string) (count int64, err error) {
	err = r.stmt.countHoroscopeById.Get(&count, horoscopeId)
	if err != nil {
		log.Println("error while count horoscope by id ", err)
	}

	return
}
