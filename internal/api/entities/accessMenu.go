package entities

type AccessMenu struct {
	AccessId string `db:"access_id" json:"access_id"`
	MenuId   string `db:"menu_id" json:"menu_id"`
	MenuName string `db:"menu_name" json:"menu_name"`
	Status   string `db:"status" json:"status"`
}
