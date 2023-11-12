package submission

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
	findSubmissions      *sqlx.Stmt
	findResultSubmission *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.SubmissionRepository {
	stmts := Statement{
		findSubmissions:      datasources.Prepare(dbReader, findSubmissions),
		findResultSubmission: datasources.Prepare(dbReader, findResultSubmission),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindSubmissions() (submissions []entities.Submission, err error) {
	err = r.stmt.findSubmissions.Select(&submissions)
	if err != nil {
		log.Println("error while find submissions ", err)
	}

	return
}

func (r Repository) FindResultSubmission(submissionId string) (resultSubmission []entities.Result, err error) {
	err = r.stmt.findResultSubmission.Select(&resultSubmission, submissionId)
	if err != nil {
		log.Println("error while find result submission ", err)
	}

	return
}
