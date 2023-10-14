package shioPoint

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
	findShioPoint   *sqlx.Stmt
	updateShioPoint *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ShioPointRepository {
	stmts := Statement{
		findShioPoint:   datasources.Prepare(dbReader, findShioPoint),
		updateShioPoint: datasources.PrepareNamed(dbWriter, updateShioPoint),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindShioPoint(elementId string) (shioPoint []entities.ShioPoint, err error) {
	err = r.stmt.findShioPoint.Select(&shioPoint, elementId)
	if err != nil {
		log.Println("error while find shio point ", err)
	}

	return
}

func (r Repository) UpdateShioPoint(params entities.RequestShioPoint) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	for index, shioId := range params.ShioId {
		data := map[string]interface{}{
			"point":       params.Point[index],
			"id_shio":     shioId,
			"id_category": params.ElementId,
		}

		_, err = tx.NamedStmt(r.stmt.updateShioPoint).Exec(data)
		if err != nil {
			log.Println("error while update shio point ", err)
		}
	}

	return
}
