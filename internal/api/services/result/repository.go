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
	findResult    *sqlx.Stmt
	findAllResult *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ResultRepository {
	stmts := Statement{
		findResult:    datasources.Prepare(dbReader, findResult),
		findAllResult: datasources.Prepare(dbReader, findAllResult),
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

func (r Repository) FindAllResult(userId string) (results []entities.Result, err error) {
	err = r.stmt.findAllResult.Select(&results, userId)
	if err != nil {
		log.Println("error while find all result ", err)
	}

	return
}
