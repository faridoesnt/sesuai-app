package pointAnswer

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"Sesuai/pkg/asql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findPointAnswer   *sqlx.Stmt
	updatePointAnswer *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.PointAnswerRepository {
	stmts := Statement{
		findPointAnswer:   datasources.Prepare(dbReader, findPointAnswer),
		updatePointAnswer: datasources.PrepareNamed(dbWriter, updatePointAnswer),
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

func (r Repository) UpdatePointAnswer(params entities.RequestPointAnswer) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	for index, pointAnswerId := range params.PointAnswerId {
		data := map[string]interface{}{
			"point":     params.Point[index],
			"id_answer": pointAnswerId,
		}

		_, err = tx.NamedStmt(r.stmt.updatePointAnswer).Exec(data)
		if err != nil {
			log.Println("error while update point answer ", err)
		}
	}

	return
}
