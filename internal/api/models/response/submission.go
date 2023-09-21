package response

type Submission struct {
	SubmissionId  string `json:"submission_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Timer         string `json:"timer"`
	TotalQuestion string `json:"total_question"`
	CreatedAt     string `json:"created_at"`
}

type ResultSubmission struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Point        string `json:"point"`
}
