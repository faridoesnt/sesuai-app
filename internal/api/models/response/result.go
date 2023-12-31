package response

type Result struct {
	SubmissionId  string `json:"submission_id"`
	Timer         string `json:"timer"`
	TotalQuestion string `json:"total_question"`
	CreatedAt     string `json:"created_at"`
	Token         string `json:"token"`
}
