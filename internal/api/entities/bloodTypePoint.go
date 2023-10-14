package entities

type BloodTypePoint struct {
	BloodTypeId   string `db:"blood_type_id" json:"blood_type_id"`
	ElementId     string `db:"category_id" json:"element_id"`
	BloodTypeName string `db:"blood_type_name" json:"blood_type_name"`
	Point         string `db:"point" json:"point"`
}

type RequestBloodTypePoint struct {
	BloodTypeId []string `json:"blood_type_id"`
	ElementId   string   `json:"element_id"`
	Point       []string `json:"point"`
}
