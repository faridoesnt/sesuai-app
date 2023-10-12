package menu

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
	findMenuIdByName *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.MenuRepository {
	stmts := Statement{
		findMenuIdByName: datasources.Prepare(dbReader, findMenuIdByName),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindMenuIdByName(name string) (menu entities.Menu, err error) {
	err = r.stmt.findMenuIdByName.Get(&menu, name)
	if err != nil {
		log.Println("error while find menu id by name ", err)
	}

	return
}
