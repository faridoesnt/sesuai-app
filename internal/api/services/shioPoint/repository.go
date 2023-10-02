package shioPoint

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
	findShioPoint *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ShioPointRepository {
	stmts := Statement{
		findShioPoint: datasources.Prepare(dbReader, findShioPoint),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindShioPoint(categoryId string) (shioPoint []entities.ShioPoint, err error) {
	err = r.stmt.findShioPoint.Select(&shioPoint, categoryId)
	if err != nil {
		log.Println("error while find shio point ", err)
	}

	return
}
