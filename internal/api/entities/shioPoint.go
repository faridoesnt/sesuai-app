package entities

type ShioPoint struct {
	ShioId     string `db:"shio_id" json:"shio_id"`
	CategoryId string `db:"category_id" json:"category_id"`
	ShioName   string `db:"shio_name" json:"shio_name"`
	Point      string `db:"point" json:"point"`
}
