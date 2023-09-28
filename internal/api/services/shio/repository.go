package shio

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
	findShio *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ShioRepository {
	stmts := Statement{
		findShio: datasources.Prepare(dbReader, findShio),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindShio() (shio []entities.Shio, err error) {
	err = r.stmt.findShio.Select(&shio)
	if err != nil {
		log.Println("error while find all shio ", err)
	}

	return
}
