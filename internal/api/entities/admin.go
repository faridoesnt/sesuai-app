package entities

type Admin struct {
	AdminId  string `db:"id_admin"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
