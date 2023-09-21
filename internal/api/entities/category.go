package entities

type Category struct {
	Id    string `db:"id_category"`
	Name  string `db:"name"`
	Photo string `db:"photo"`
}

type RequestCategory struct {
	Name    string `db:"name" json:"name"`
	Photo   string `db:"photo" json:"photo"`
	AdminId string `db:"admin_id"`
}
