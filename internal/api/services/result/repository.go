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
	findResultBySubmissionId *sqlx.Stmt
	findAllResult            *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ResultRepository {
	stmts := Statement{
		findResultBySubmissionId: datasources.Prepare(dbReader, findResultBySubmissionId),
		findAllResult:            datasources.Prepare(dbReader, findAllResult),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error) {
	err = r.stmt.findResultBySubmissionId.Select(&results, userId, submissionId)
	if err != nil {
		log.Println("error while find result by submission id ", err)
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
