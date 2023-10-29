package result

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
	findResult *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ResultRepository {
	stmts := Statement{
		findResult: datasources.Prepare(dbReader, findResult),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindResult(userId string) (results []entities.Result, err error) {
	err = r.stmt.findResult.Select(&results, userId)
	if err != nil {
		log.Println("error while find result ", err)
	}

	return
}
