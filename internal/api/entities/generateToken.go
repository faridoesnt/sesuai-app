package entities

type GenerateToken struct {
	Id        string `db:"id_token"`
	Admin     string `db:"admin"`
	Token     string `db:"token"`
	Status    string `db:"status"`
	CreatedAt string `db:"created_at"`
}

type UseToken struct {
	Id string `db:"id_token" json:"id_token"`
}
