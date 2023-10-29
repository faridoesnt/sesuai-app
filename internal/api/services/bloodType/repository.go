package bloodType

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
	findBloodType     *sqlx.Stmt
	findBloodTypeUser *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.BloodTypeRepository {
	stmts := Statement{
		findBloodType:     datasources.Prepare(dbReader, findBloodType),
		findBloodTypeUser: datasources.Prepare(dbReader, findBloodTypeUser),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindBloodType() (bloodType []entities.BloodType, err error) {
	err = r.stmt.findBloodType.Select(&bloodType)
	if err != nil {
		log.Println("error while find blood type ", err)
	}

	return
}

func (r Repository) FindBloodTypeUser(userId string) (bloodType entities.BloodType, err error) {
	err = r.stmt.findBloodTypeUser.Get(&bloodType, userId)
	if err != nil {
		log.Println("error while find blood type user ", err)
	}

	return
}
