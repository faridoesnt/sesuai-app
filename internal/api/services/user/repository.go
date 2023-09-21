package user

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
	findUserByEmail  *sqlx.Stmt
	refreshToken     *sqlx.Stmt
	insertUser       *sqlx.NamedStmt
	countPhoneNumber *sqlx.Stmt
	findUserLoggedIn *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.UserRepository {
	stmts := Statement{
		findUserByEmail:  datasources.Prepare(dbReader, findUserByEmail),
		refreshToken:     datasources.Prepare(dbWriter, refreshToken),
		insertUser:       datasources.PrepareNamed(dbWriter, insertUser),
		countPhoneNumber: datasources.Prepare(dbReader, countPhoneNumber),
		findUserLoggedIn: datasources.Prepare(dbReader, findUserLoggedIn),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) FindUserByEmail(email string) (user entities.User, err error) {
	err = r.stmt.findUserByEmail.Get(&user, email)
	if err != nil {
		log.Println("error while find user ", err)
	}

	return
}

func (r Repository) RefreshToken(email, token string) (err error) {
	_, err = r.stmt.refreshToken.Exec(token, email)
	if err != nil {
		log.Println("error while refresh token user ", err)
	}

	return
}

func (r Repository) InsertUser(user entities.RequestRegister) (err error) {
	_, err = r.stmt.insertUser.Exec(user)
	if err != nil {
		log.Println("error while register ", err)
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

func (r Repository) FindUserLoggedIn(userId, token string) (user entities.User, err error) {
	err = r.stmt.findUserLoggedIn.Get(&user, userId, token)
	if err != nil {
		log.Println("error while find user logged in ", err)
	}

	return
}
