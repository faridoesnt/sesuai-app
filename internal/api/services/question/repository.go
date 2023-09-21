package question

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findQuestion   *sqlx.Stmt
	insertQuestion *sqlx.NamedStmt
	deleteQuestion *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.QuestionRepository {
	stmts := Statement{
		findQuestion:   datasources.Prepare(dbReader, findQuestion),
		insertQuestion: datasources.PrepareNamed(dbWriter, insertQuestion),
		deleteQuestion: datasources.Prepare(dbWriter, deleteQuestion),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindQuestions(category string) (questions []entities.Question, err error) {
	query := `
		SELECT
			q.id_question,
			c.id_category,
			c.name as category,
			c.photo,
			q.question_ina,
			q.question_eng
		FROM
		    question q
		LEFT JOIN category c
			ON q.id_category = c.id_category
	`

	if category != "" {
		categoryLike := fmt.Sprintf("'%s%%'", category)
		query += fmt.Sprintf(`
			WHERE c.name LIKE %s
		`, categoryLike)
	}

	err = r.dbReader.Select(&questions, query)
	if err != nil {
		log.Println("error while find questions ", err)
	}

	return
}

func (r Repository) FindQuestion(categoryId string) (question entities.Question, err error) {
	err = r.stmt.findQuestion.Get(&question, categoryId)
	if err != nil {
		log.Println("error while find question ", err)
	}

	return
}

func (r Repository) InsertQuestion(params entities.RequestQuestion) (err error) {
	_, err = r.stmt.insertQuestion.Exec(params)
	if err != nil {
		log.Println("error while insert question ", err)
	}

	return
}

func (r Repository) DeleteQuestion(questionId string) (err error) {
	_, err = r.stmt.deleteQuestion.Exec(questionId)
	if err != nil {
		log.Println("error while delete question ", err)
	}

	return
}
