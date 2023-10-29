package generateToken

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
	findGenerateToken        *sqlx.Stmt
	insertNewToken           *sqlx.Stmt
	updateToken              *sqlx.Stmt
	findGenerateTokenByToken *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.GenerateTokenRepository {
	stmts := Statement{
		findGenerateToken:        datasources.Prepare(dbReader, findGenerateToken),
		insertNewToken:           datasources.Prepare(dbWriter, insertNewToken),
		updateToken:              datasources.Prepare(dbWriter, updateToken),
		findGenerateTokenByToken: datasources.Prepare(dbReader, findGenerateTokenByToken),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindGenerateToken(adminId string) (tokens []entities.GenerateToken, err error) {
	err = r.stmt.findGenerateToken.Select(&tokens, adminId)
	if err != nil {
		log.Println("error while find generate token ", err)
	}

	return
}

func (r Repository) InsertNewToken(adminId, token string) (err error) {
	_, err = r.stmt.insertNewToken.Exec(adminId, token)
	if err != nil {
		log.Println("error while insert new token ", err)
	}

	return
}

func (r Repository) UpdateToken(tokenId string) (err error) {
	_, err = r.stmt.updateToken.Exec(tokenId)
	if err != nil {
		log.Println("error while update token ", err)
	}

	return
}

func (r Repository) FindGenerateTokenByToken(params string) (token entities.GenerateToken, err error) {
	err = r.stmt.findGenerateTokenByToken.Get(&token, params)
	if err != nil {
		log.Println("error while find generate token by token ", err)
	}

	return
}
