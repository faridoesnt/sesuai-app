package admin

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
	findAdminByEmail  *sqlx.Stmt
	refreshToken      *sqlx.Stmt
	findAdminLoggedIn *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.AdminRepository {
	stmts := Statement{
		findAdminByEmail:  datasources.Prepare(dbReader, findAdminByEmail),
		refreshToken:      datasources.Prepare(dbWriter, refreshToken),
		findAdminLoggedIn: datasources.Prepare(dbReader, findAdminLoggedIn),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindAdminByEmail(email string) (admin entities.Admin, err error) {
	err = r.stmt.findAdminByEmail.Get(&admin, email)
	if err != nil {
		log.Println("error while find admin ", err)
	}

	return
}

func (r Repository) RefreshToken(email, token string) (err error) {
	_, err = r.stmt.refreshToken.Exec(token, email)
	if err != nil {
		log.Println("error while refresh token admin ", err)
	}

	return
}

func (r Repository) FindAdminLoggedIn(adminId, token string) (admin entities.Admin, err error) {
	err = r.stmt.findAdminLoggedIn.Get(&admin, adminId, token)
	if err != nil {
		log.Println("error while find admin logged in ", err)
	}

	return
}
