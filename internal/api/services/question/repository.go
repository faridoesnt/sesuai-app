package question

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
	findQuestionsByCategoryId    *sqlx.Stmt
	findAllQuestionsByCategoryId *sqlx.Stmt
	findQuestion                 *sqlx.Stmt
	insertQuestion               *sqlx.NamedStmt
	updateQuestion               *sqlx.NamedStmt
	deleteQuestion               *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.QuestionRepository {
	stmts := Statement{
		findQuestionsByCategoryId:    datasources.Prepare(dbReader, findQuestionsByCategoryId),
		findAllQuestionsByCategoryId: datasources.Prepare(dbReader, findAllQuestionsByCategoryId),
		findQuestion:                 datasources.Prepare(dbReader, findQuestion),
		insertQuestion:               datasources.PrepareNamed(dbWriter, insertQuestion),
		updateQuestion:               datasources.PrepareNamed(dbWriter, updateQuestion),
		deleteQuestion:               datasources.Prepare(dbWriter, deleteQuestion),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindQuestionsByCategoryId(categoryId string) (questions []entities.Question, err error) {
	err = r.stmt.findQuestionsByCategoryId.Select(&questions, categoryId)
	if err != nil {
		log.Println("error while find questions ", err)
	}

	return
}

func (r Repository) FindAllQuestionsByCategoryId(categoryId string) (questions []entities.Question, err error) {
	err = r.stmt.findAllQuestionsByCategoryId.Select(&questions, categoryId)
	if err != nil {
		log.Println("error while find all questions by category id ", err)
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

func (r Repository) UpdateQuestion(questionId string, params entities.RequestQuestion) (err error) {
	data := map[string]interface{}{
		"id_category":  params.CategoryId,
		"question_ina": params.QuestionIna,
		"question_eng": params.QuestionEn,
		"id_question":  questionId,
	}

	_, err = r.stmt.updateQuestion.Exec(data)
	if err != nil {
		log.Println("error while update question ", err)
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
