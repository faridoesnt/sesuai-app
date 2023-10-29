package entities

type Result struct {
	ElementName string `db:"element_name" json:"element_name"`
	Point       string `db:"point" json:"point"`
}

type RequestAllResult struct {
	Token string `db:"token" json:"token"`
}
