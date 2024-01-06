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
	findResults                 *sqlx.Stmt
	findResultBySubmissionId    *sqlx.Stmt
	findAllResultBySubmissionId *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ResultRepository {
	stmts := Statement{
		findResults:                 datasources.Prepare(dbReader, findResults),
		findResultBySubmissionId:    datasources.Prepare(dbReader, findResultBySubmissionId),
		findAllResultBySubmissionId: datasources.Prepare(dbReader, findAllResultBySubmissionId),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindResults(userId string) (results []entities.Submission, err error) {
	err = r.stmt.findResults.Select(&results, userId)
	if err != nil {
		log.Println("error while find results ", err)
	}

	return
}

func (r Repository) FindResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error) {
	err = r.stmt.findResultBySubmissionId.Select(&results, userId, submissionId)
	if err != nil {
		log.Println("error while find result by submission id ", err)
	}

	return
}

func (r Repository) FindAllResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error) {
	err = r.stmt.findAllResultBySubmissionId.Select(&results, userId, submissionId)
	if err != nil {
		log.Println("error while find all result by submission id ", err)
	}

	return
}
