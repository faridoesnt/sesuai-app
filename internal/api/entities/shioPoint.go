package entities

type ShioPoint struct {
	ShioId    string `db:"shio_id" json:"shio_id"`
	ElementId string `db:"category_id" json:"element_id"`
	ShioName  string `db:"shio_name" json:"shio_name"`
	Point     string `db:"point" json:"point"`
}

type RequestShioPoint struct {
	ShioId    []string `json:"shio_id"`
	ElementId string   `json:"element_id"`
	Point     []string `json:"point"`
}
