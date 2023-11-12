package entities

type Submission struct {
	SubmissionId    string `db:"submission_id"`
	Name            string `db:"name"`
	Email           string `db:"email"`
	Timer           string `db:"timer"`
	TotalSubmission string `db:"total_submission"`
	TotalQuestion   string `db:"total_question"`
	CreatedAt       string `db:"created_at"`
	Token           string `db:"token"`
}

type ResultSubmission struct {
	ElementId   string `db:"category_id"`
	ElementName string `db:"category_name"`
	Point       string `db:"point"`
}
