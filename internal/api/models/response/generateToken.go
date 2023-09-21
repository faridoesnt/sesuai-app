package response

type GenerateToken struct {
	Id        string `json:"id_token"`
	Admin     string `json:"admin"`
	Token     string `json:"token"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
