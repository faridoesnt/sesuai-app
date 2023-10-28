package questionTest

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
	findQuestionsTest *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.QuestionTestRepository {
	stmts := Statement{
		findQuestionsTest: datasources.Prepare(dbReader, findQuestionsTest),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindQuestionsTest() (questionsTest []entities.QuestionTest, err error) {
	err = r.stmt.findQuestionsTest.Select(&questionsTest)
	if err != nil {
		log.Println("error while find questions test ", err)
	}

	return
}
