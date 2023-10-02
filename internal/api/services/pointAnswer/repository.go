package pointAnswer

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
	findPointAnswer *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.PointAnswerRepository {
	stmts := Statement{
		findPointAnswer: datasources.Prepare(dbReader, findPointAnswer),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindPointAnswer() (pointAnswer []entities.PointAnswer, err error) {
	err = r.stmt.findPointAnswer.Select(&pointAnswer)
	if err != nil {
		log.Println("error while find point answer ", err)
	}

	return
}
