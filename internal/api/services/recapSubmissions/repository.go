package recapSubmissions

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
	countRecapSubmissionsUser *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.RecapSubmissionsRepository {
	stmts := Statement{
		countRecapSubmissionsUser: datasources.Prepare(dbReader, countRecapSubmissionsUser),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindRecapUser(params entities.RequestRecapSubmissions) (recapUser []entities.RecapUser, err error) {
	var existWhere int
	var value []interface{}

	query := findRecapSubmissions

	if params.Horoscope != "all" {
		query += " WHERE" + whereHoroscope
		value = append(value, params.Horoscope)

		existWhere++
	}

	if params.Shio != "all" {
		if existWhere > 0 {
			query += " AND" + whereShio
		} else {
			query += " WHERE" + whereShio
		}

		value = append(value, params.Shio)

		existWhere++
	}

	if params.BloodType != "all" {
		if existWhere > 0 {
			query += " AND" + whereBloodType
		} else {
			query += " WHERE" + whereBloodType
		}

		value = append(value, params.BloodType)

		existWhere++
	}

	if params.Gender != "all" {
		if existWhere > 0 {
			query += " AND" + whereGender
		} else {
			query += " WHERE" + whereGender
		}

		value = append(value, params.Gender)

		existWhere++
	}

	query += orderBy

	err = r.dbReader.Select(&recapUser, query, value...)
	if err != nil {
		log.Println("error while find recap user ", err)
	}

	return
}

func (r Repository) CountRecapSubmissionsUser(userId string) (recapSubmissions entities.RecapSubmissions, err error) {
	err = r.stmt.countRecapSubmissionsUser.Get(&recapSubmissions, userId)
	if err != nil {
		log.Println("error while count recap submissions user ", err)
	}

	return
}
