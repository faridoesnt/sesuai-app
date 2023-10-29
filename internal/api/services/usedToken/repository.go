package usedToken

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	insertUsedToken *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.UsedTokenRepository {
	stmts := Statement{
		insertUsedToken: datasources.Prepare(dbWriter, insertUsedToken),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) InsertUsedToken(tokenId, userId string) (err error) {
	_, err = r.stmt.insertUsedToken.Exec(tokenId, userId)
	if err != nil {
		log.Println("error while insert used token ", err)
	}

	return
}
