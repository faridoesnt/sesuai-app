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
	findAdmins        *sqlx.Stmt
	findAdminByEmail  *sqlx.Stmt
	refreshToken      *sqlx.Stmt
	findAdminLoggedIn *sqlx.Stmt
	countEmail        *sqlx.Stmt
	countPhoneNumber  *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.AdminRepository {
	stmts := Statement{
		findAdmins:        datasources.Prepare(dbReader, findAdmins),
		findAdminByEmail:  datasources.Prepare(dbReader, findAdminByEmail),
		refreshToken:      datasources.Prepare(dbWriter, refreshToken),
		findAdminLoggedIn: datasources.Prepare(dbReader, findAdminLoggedIn),
		countEmail:        datasources.Prepare(dbReader, countEmail),
		countPhoneNumber:  datasources.Prepare(dbReader, countPhoneNumber),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindAdmins() (admins []entities.AdminList, err error) {
	err = r.stmt.findAdmins.Select(&admins)
	if err != nil {
		log.Println("error while find admins ", err)
	}

	return
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

func (r Repository) CountEmail(email string) (total int64, err error) {
	err = r.stmt.countEmail.Get(&total, email)
	if err != nil {
		log.Println("error while count email ", err)
	}

	return
}

func (r Repository) CountPhoneNumber(phoneNumber string) (total int64, err error) {
	err = r.stmt.countPhoneNumber.Get(&total, phoneNumber)
	if err != nil {
		log.Println("error while count phone number ", err)
	}

	return
}
