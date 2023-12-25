package entities

type Admin struct {
	AdminId      string `db:"id_admin"`
	FullName     string `db:"full_name"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	IsSuperAdmin bool   `db:"is_super_admin"`
}

type AdminList struct {
	AdminId     string   `db:"admin_id" json:"admin_id"`
	FullName    string   `db:"full_name" json:"full_name"`
	Email       string   `db:"email" json:"email"`
	PhoneNumber string   `db:"phone_number" json:"phone_number"`
	AccessMenu  []string `json:"access_menu"`
}

type RequestAdmin struct {
	FullName    string   `db:"full_name" json:"full_name"`
	Email       string   `db:"email" json:"email"`
	PhoneNumber string   `db:"phone_number" json:"phone_number"`
	Password    string   `db:"password" json:"password"`
	Access      []string `db:"access" json:"access"`
}
