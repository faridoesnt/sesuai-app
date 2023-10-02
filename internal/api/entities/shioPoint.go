package entities

type ShioPoint struct {
	ShioId     string `db:"shio_id" json:"shio_id"`
	CategoryId string `db:"category_id" json:"category_id"`
	ShioName   string `db:"shio_name" json:"shio_name"`
	Point      string `db:"point" json:"point"`
}

type RequestShioPoint struct {
	ShioId     []string `json:"shio_id"`
	CategoryId string   `json:"category_id"`
	Point      []string `json:"point"`
}
