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
	insertUsedToken       *sqlx.Stmt
	countUserToken        *sqlx.Stmt
	findUsedTokenByUserId *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.UsedTokenRepository {
	stmts := Statement{
		insertUsedToken:       datasources.Prepare(dbWriter, insertUsedToken),
		countUserToken:        datasources.Prepare(dbReader, countUserToken),
		findUsedTokenByUserId: datasources.Prepare(dbReader, findUsedTokenByUserId),
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

func (r Repository) CountUserToken(token, userId string) (total int64, err error) {
	err = r.stmt.countUserToken.Get(&total, token, userId)
	if err != nil {
		log.Println("error while count user token ", err)
	}

	return
}

func (r Repository) FindUsedTokenByUserId(userId string) (token string, err error) {
	err = r.stmt.findUsedTokenByUserId.Get(&token, userId)
	if err != nil {
		log.Println("error while find used token by user id ", err)
	}

	return
}
