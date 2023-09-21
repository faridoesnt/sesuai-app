package constracts

import (
	"github.com/jmoiron/sqlx"
)

type Datasources struct {
	WriterDB *sqlx.DB `json:"writer-db"`
	ReaderDB *sqlx.DB `json:"reader-db"`
}
