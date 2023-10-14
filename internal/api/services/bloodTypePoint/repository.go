package bloodTypePoint

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
	findBloodTypePoint   *sqlx.Stmt
	updateBloodTypePoint *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.BloodTypePointRepository {
	stmts := Statement{
		findBloodTypePoint:   datasources.Prepare(dbReader, findBloodTypePoint),
		updateBloodTypePoint: datasources.PrepareNamed(dbWriter, updateBloodTypePoint),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindBloodTypePoint(elementId string) (bloodTypePoint []entities.BloodTypePoint, err error) {
	err = r.stmt.findBloodTypePoint.Select(&bloodTypePoint, elementId)
	if err != nil {
		log.Println("error while find blood type point ", err)
	}

	return
}

func (r Repository) UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	for index, bloodTypeId := range params.BloodTypeId {
		data := map[string]interface{}{
			"point":         params.Point[index],
			"id_blood_type": bloodTypeId,
			"id_category":   params.ElementId,
		}

		_, err = tx.NamedStmt(r.stmt.updateBloodTypePoint).Exec(data)
		if err != nil {
			log.Println("error while update horoscope point ", err)
		}
	}

	return
}
