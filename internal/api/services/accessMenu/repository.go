package accessMenu

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
	findAccessMenuByAdminId *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.AccessMenuRepository {
	stmts := Statement{
		findAccessMenuByAdminId: datasources.Prepare(dbReader, findAccessMenuByAdminId),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindAccessMenuByAdminId(adminId string) (accessMenus []entities.AccessMenu, err error) {
	err = r.stmt.findAccessMenuByAdminId.Select(&accessMenus, adminId)
	if err != nil {
		log.Println("error while find access menu by admin id ", err)
	}

	return
}
