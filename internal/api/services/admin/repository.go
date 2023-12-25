package admin

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/datasources"
	"Sesuai/internal/api/entities"
	"Sesuai/pkg/asql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	findAdmins          *sqlx.Stmt
	findAdminById       *sqlx.Stmt
	findAdminByEmail    *sqlx.Stmt
	refreshToken        *sqlx.Stmt
	findAdminLoggedIn   *sqlx.Stmt
	countEmail          *sqlx.Stmt
	countPhoneNumber    *sqlx.Stmt
	insertAdmin         *sqlx.NamedStmt
	insertAccessMenu    *sqlx.NamedStmt
	updateAdmin         *sqlx.NamedStmt
	deleteAccessMenu    *sqlx.NamedStmt
	countAdmin          *sqlx.Stmt
	countAdminWithToken *sqlx.Stmt
	deleteAdmin         *sqlx.NamedStmt
	changePassword      *sqlx.Stmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB) constracts.AdminRepository {
	stmts := Statement{
		findAdmins:          datasources.Prepare(dbReader, findAdmins),
		findAdminById:       datasources.Prepare(dbReader, findAdminById),
		findAdminByEmail:    datasources.Prepare(dbReader, findAdminByEmail),
		refreshToken:        datasources.Prepare(dbWriter, refreshToken),
		findAdminLoggedIn:   datasources.Prepare(dbReader, findAdminLoggedIn),
		countEmail:          datasources.Prepare(dbReader, countEmail),
		countPhoneNumber:    datasources.Prepare(dbReader, countPhoneNumber),
		insertAdmin:         datasources.PrepareNamed(dbWriter, insertAdmin),
		insertAccessMenu:    datasources.PrepareNamed(dbWriter, insertAccessMenu),
		updateAdmin:         datasources.PrepareNamed(dbWriter, updateAdmin),
		deleteAccessMenu:    datasources.PrepareNamed(dbWriter, deleteAccessMenu),
		countAdmin:          datasources.Prepare(dbReader, countAdmin),
		countAdminWithToken: datasources.Prepare(dbReader, countAdminWithToken),
		deleteAdmin:         datasources.PrepareNamed(dbWriter, deleteAdmin),
		changePassword:      datasources.Prepare(dbWriter, changePassword),
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

func (r Repository) FindAdminById(adminId string) (admin entities.AdminList, err error) {
	err = r.stmt.findAdminById.Get(&admin, adminId)
	if err != nil {
		log.Println("error while find admin by id ", err)
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

func (r Repository) InsertAdmin(params entities.RequestAdmin) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	data := map[string]interface{}{
		"fullname":       params.FullName,
		"email":          params.Email,
		"password":       params.Password,
		"phone_number":   params.PhoneNumber,
		"is_super_admin": 0,
	}

	res, err := tx.NamedStmt(r.stmt.insertAdmin).Exec(data)
	if err != nil {
		log.Println("error while insert admin ", err)
	}

	resId, _ := res.LastInsertId()

	adminId := strconv.FormatInt(resId, 10)

	for _, val := range params.Access {
		data = map[string]interface{}{
			"id_menu":  val,
			"id_admin": adminId,
			"status":   "write",
		}

		_, err = tx.NamedStmt(r.stmt.insertAccessMenu).Exec(data)
		if err != nil {
			log.Println("error while insert access menu admin ", err)
		}
	}

	return
}

func (r Repository) UpdateAdmin(adminId string, params entities.RequestAdmin) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	data := map[string]interface{}{
		"fullname":     params.FullName,
		"email":        params.Email,
		"phone_number": params.PhoneNumber,
		"id_admin":     adminId,
	}

	_, err = tx.NamedStmt(r.stmt.updateAdmin).Exec(data)
	if err != nil {
		log.Println("error while update admin ", err)
	}

	_, err = tx.NamedStmt(r.stmt.deleteAccessMenu).Exec(data)
	if err != nil {
		log.Println("error while delete access menu admin ", err)
	}

	for _, val := range params.Access {
		data = map[string]interface{}{
			"id_menu":  val,
			"id_admin": adminId,
			"status":   "write",
		}

		_, err = tx.NamedStmt(r.stmt.insertAccessMenu).Exec(data)
		if err != nil {
			log.Println("error while update insert menu admin ", err)
		}
	}

	return
}

func (r Repository) CountAdmin(adminId string) (total int64, err error) {
	err = r.stmt.countAdmin.Get(&total, adminId)
	if err != nil {
		log.Println("error while count admin ", err)
	}

	return
}

func (r Repository) CountAdminWithToken(adminId, token string) (total int64, err error) {
	err = r.stmt.countAdminWithToken.Get(&total, adminId, token)
	if err != nil {
		log.Println("error while count admin with token ", err)
	}

	return
}

func (r Repository) DeleteAdmin(adminId string) (err error) {
	tx, err := r.dbWriter.Beginx()
	if err != nil {
		return err
	}

	defer asql.ReleaseTx(tx, &err)

	data := map[string]interface{}{
		"id_admin": adminId,
	}

	_, err = tx.NamedStmt(r.stmt.deleteAccessMenu).Exec(data)
	if err != nil {
		log.Println("error while delete access menu admin ", err)
	}

	_, err = tx.NamedStmt(r.stmt.deleteAdmin).Exec(data)
	if err != nil {
		log.Println("error while delete admin ", err)
	}

	return
}

func (r Repository) ChangePassword(adminId, newPassword string) (err error) {
	_, err = r.stmt.changePassword.Exec(newPassword, adminId)
	if err != nil {
		log.Println("error while change password ", err)
	}

	return
}
