package entities

type Result struct {
	Id           string `db:"id" json:"id"`
	ElementName  string `db:"element_name" json:"element_name"`
	ElementImage string `db:"element_image" json:"element_image"`
	Point        string `db:"point" json:"point"`
	Note         string `db:"-" json:"note"`
}

type TokenResult struct {
	SubmissionId string `db:"submission_id" json:"submission_id"`
	Token        string `db:"token" json:"token"`
}
