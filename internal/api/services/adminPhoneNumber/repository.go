package adminPhoneNumber

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
	findAdminPhoneNumber *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.AdminPhoneNumberRepository {
	stmts := Statement{
		findAdminPhoneNumber: datasources.Prepare(dbReader, findAdminPhoneNumber),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindAdminPhoneNumber() (adminPhoneNumber entities.AdminPhoneNumber, err error) {
	err = r.stmt.findAdminPhoneNumber.Get(&adminPhoneNumber)
	if err != nil {
		log.Println("error while find admin phone number ", err)
	}

	return
}
