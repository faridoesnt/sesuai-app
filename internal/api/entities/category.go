package entities

type Category struct {
	Id    string `db:"id_category"`
	Name  string `db:"name"`
	Photo string `db:"photo"`
}

type RequestCategory struct {
	Name     string `db:"name" json:"name"`
	FileName string `db:"photo" json:"filename"`
	Image    string `db:"-" json:"image"`
	AdminId  string `db:"admin_id"`
}
