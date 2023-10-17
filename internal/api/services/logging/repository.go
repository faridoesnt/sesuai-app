package logging

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
	insertMobileLogAdmin *sqlx.NamedStmt
	insertMobileLogUser  *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.LoggingRepository {
	stmts := Statement{
		insertMobileLogAdmin: datasources.PrepareNamed(dbWriter, insertMobileLogAdmin),
		insertMobileLogUser:  datasources.PrepareNamed(dbWriter, insertMobileLogUser),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) InsertMobileLogAdmin(mobileLog entities.MobileLogAdmin) (err error) {
	_, err = r.stmt.insertMobileLogAdmin.Exec(mobileLog)
	if err != nil {
		log.Println("error while insert mobile log admin")
	}

	return
}

func (r Repository) InsertMobileLogUser(mobileLog entities.MobileLogUser) (err error) {
	_, err = r.stmt.insertMobileLogUser.Exec(mobileLog)
	if err != nil {
		log.Println("error while insert mobile log user")
	}

	return
}
