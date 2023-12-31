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
	findUserByEmail             *sqlx.Stmt
	refreshToken                *sqlx.Stmt
	insertUser                  *sqlx.NamedStmt
	countPhoneNumber            *sqlx.Stmt
	findUserLoggedIn            *sqlx.Stmt
	findProfileUser             *sqlx.Stmt
	updateProfileUser           *sqlx.NamedStmt
	countEmailAlreadyUsed       *sqlx.Stmt
	countPhoneNumberAlreadyUsed *sqlx.Stmt
	findUserById                *sqlx.Stmt
	changePassword              *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.UserRepository {
	stmts := Statement{
		findUserByEmail:             datasources.Prepare(dbReader, findUserByEmail),
		refreshToken:                datasources.Prepare(dbWriter, refreshToken),
		insertUser:                  datasources.PrepareNamed(dbWriter, insertUser),
		countPhoneNumber:            datasources.Prepare(dbReader, countPhoneNumber),
		findUserLoggedIn:            datasources.Prepare(dbReader, findUserLoggedIn),
		findProfileUser:             datasources.Prepare(dbReader, findProfileUser),
		updateProfileUser:           datasources.PrepareNamed(dbWriter, updateProfileUser),
		countEmailAlreadyUsed:       datasources.Prepare(dbReader, countEmailAlreadyUsed),
		countPhoneNumberAlreadyUsed: datasources.Prepare(dbReader, countPhoneNumberAlreadyUsed),
		findUserById:                datasources.Prepare(dbReader, findUserById),
		changePassword:              datasources.Prepare(dbWriter, changePassword),
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

func (r Repository) FindProfileUser(userId string) (profileUser entities.User, err error) {
	err = r.stmt.findProfileUser.Get(&profileUser, userId)
	if err != nil {
		log.Println("error while find profile user ", err)
	}

	return
}

func (r Repository) UpdateProfileUser(params entities.UpdateProfile) (err error) {
	_, err = r.stmt.updateProfileUser.Exec(params)
	if err != nil {
		log.Println("error while update profile user ", err)
	}

	return
}

func (r Repository) CountEmailAlreadyUsed(email, userId string) (total int64, err error) {
	err = r.stmt.countEmailAlreadyUsed.Get(&total, email, userId)
	if err != nil {
		log.Println("error while count email already used ", err)
	}

	return
}

func (r Repository) CountPhoneNumberAlreadyUsed(phoneNumber, userId string) (total int64, err error) {
	err = r.stmt.countPhoneNumberAlreadyUsed.Get(&total, phoneNumber, userId)
	if err != nil {
		log.Println("error while count phone number already used ", err)
	}

	return
}

func (r Repository) FindUserById(userId string) (user entities.User, err error) {
	err = r.stmt.findUserById.Get(&user, userId)
	if err != nil {
		log.Println("error while find user by id ", err)
	}

	return
}

func (r Repository) ChangePassword(userId, newPassword string) (err error) {
	_, err = r.stmt.changePassword.Exec(newPassword, userId)
	if err != nil {
		log.Println("error while change password ", err)
	}

	return
}
