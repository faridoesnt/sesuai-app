package shio

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
	findShio            *sqlx.Stmt
	findShioByShioName  *sqlx.Stmt
	findShioUser        *sqlx.Stmt
	findShioSupportUser *sqlx.Stmt
	countShioById       *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.ShioRepository {
	stmts := Statement{
		findShio:            datasources.Prepare(dbReader, findShio),
		findShioByShioName:  datasources.Prepare(dbReader, findShioByShioName),
		findShioUser:        datasources.Prepare(dbReader, findShioUser),
		findShioSupportUser: datasources.Prepare(dbReader, findShioSupportUser),
		countShioById:       datasources.Prepare(dbReader, countShioById),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindShio() (shio []entities.Shio, err error) {
	err = r.stmt.findShio.Select(&shio)
	if err != nil {
		log.Println("error while find all shio ", err)
	}

	return
}

func (r Repository) FindShioByShioName(shioName string) (shio entities.Shio, err error) {
	err = r.stmt.findShioByShioName.Get(&shio, shioName)
	if err != nil {
		log.Println("error while find shio by shio name ", err)
	}

	return
}

func (r Repository) FindShioUser(userId string) (shio entities.Shio, err error) {
	err = r.stmt.findShioUser.Get(&shio, userId)
	if err != nil {
		log.Println("error while find shio user ", err)
	}

	return
}

func (r Repository) FindShioSupportUser(userId string) (shio entities.Shio, err error) {
	err = r.stmt.findShioSupportUser.Get(&shio, userId)
	if err != nil {
		log.Println("error while find shio support user ", err)
	}

	return
}

func (r Repository) CountShioById(shioId string) (count int64, err error) {
	err = r.stmt.countShioById.Get(&count, shioId)
	if err != nil {
		log.Println("error while count shio by id ", err)
	}

	return
}
