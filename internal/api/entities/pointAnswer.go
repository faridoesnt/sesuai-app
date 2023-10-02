package entities

type PointAnswer struct {
	Id    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Point string `db:"point" json:"point"`
}
