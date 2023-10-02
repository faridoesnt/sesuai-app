package bloodTypePoint

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
	findBloodTypePoint *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.BloodTypePointRepository {
	stmts := Statement{
		findBloodTypePoint: datasources.Prepare(dbReader, findBloodTypePoint),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindBloodTypePoint(categoryId string) (bloodTypePoint []entities.BloodTypePoint, err error) {
	err = r.stmt.findBloodTypePoint.Select(&bloodTypePoint, categoryId)
	if err != nil {
		log.Println("error while find blood type point ", err)
	}

	return
}
