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
	findSubmissions            *sqlx.Stmt
	findSubmissionsByEmailUser *sqlx.Stmt
	findSubmissionsByFullName  *sqlx.Stmt
	findResultSubmission       *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.SubmissionRepository {
	stmts := Statement{
		findSubmissions:            datasources.Prepare(dbReader, findSubmissions),
		findSubmissionsByEmailUser: datasources.Prepare(dbReader, findSubmissionsByEmailUser),
		findSubmissionsByFullName:  datasources.Prepare(dbReader, findSubmissionsByFullName),
		findResultSubmission:       datasources.Prepare(dbReader, findResultSubmission),
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

func (r Repository) FindSubmissionsByEmailUser(emailUser string) (submissions []entities.Submission, err error) {
	err = r.stmt.findSubmissionsByEmailUser.Select(&submissions, emailUser)
	if err != nil {
		log.Println("error while find submissions by email user ", err)
	}

	return
}

func (r Repository) FindSubmissionsByFullName(fullName string) (submissions []entities.Submission, err error) {
	err = r.stmt.findSubmissionsByFullName.Select(&submissions, "%"+fullName+"%")
	if err != nil {
		log.Println("error while find submissions by full name ", err)
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
