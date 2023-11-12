package response

type Submission struct {
	SubmissionId  string `json:"submission_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Timer         string `json:"timer"`
	TotalQuestion string `json:"total_question"`
	CreatedAt     string `json:"created_at"`
	Token         string `json:"token"`
}
