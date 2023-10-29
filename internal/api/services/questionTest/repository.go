package questionTest

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"Sesuai/pkg/asql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findQuestionsTest     *sqlx.Stmt
	insertSubmission      *sqlx.NamedStmt
	insertSubSubmission   *sqlx.NamedStmt
	insertPointSubmission *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.QuestionTestRepository {
	stmts := Statement{
		findQuestionsTest:     datasources.Prepare(dbReader, findQuestionsTest),
		insertSubmission:      datasources.PrepareNamed(dbWriter, insertSubmission),
		insertSubSubmission:   datasources.PrepareNamed(dbWriter, insertSubSubmission),
		insertPointSubmission: datasources.PrepareNamed(dbWriter, insertPointSubmission),
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

func (r Repository) SubmitQuestionTest(params entities.SubmitQuestionTest, userId string, totalPointQuestionsByElement map[string]float64) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	dataInsertSubmission := map[string]interface{}{
		"id_user": userId,
		"time":    params.Timer,
	}

	res, err := tx.NamedStmt(r.stmt.insertSubmission).Exec(dataInsertSubmission)
	if err != nil {
		log.Println("error while insert submission ", err)
	}

	resId, _ := res.LastInsertId()

	submissionId := strconv.FormatInt(resId, 10)

	dataInsertSubSubmission := make(map[string]interface{})

	for _, val := range params.Submit {
		dataInsertSubSubmission["id_submission"] = submissionId
		dataInsertSubSubmission["id_question"] = val.QuestionId
		dataInsertSubSubmission["id_answer"] = val.AnswerId

		_, err = tx.NamedStmt(r.stmt.insertSubSubmission).Exec(dataInsertSubSubmission)
		if err != nil {
			log.Println("error while insert sub submission ", err)
		}
	}

	dataInsertPointSubmission := make(map[string]interface{})

	for elementId, point := range totalPointQuestionsByElement {
		dataInsertPointSubmission["id_submission"] = submissionId
		dataInsertPointSubmission["id_category"] = elementId
		dataInsertPointSubmission["point"] = point

		_, err = tx.NamedStmt(r.stmt.insertPointSubmission).Exec(dataInsertPointSubmission)
		if err != nil {
			log.Println("error while insert point submission ", err)
		}
	}

	return
}
