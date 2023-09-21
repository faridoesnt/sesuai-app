package datasources

import (
	"Sesuai/internal/api/constants"
	"Sesuai/pkg/alog"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("error while preparing statement: %s", err)))
		alog.Logger.Error(errors.New(fmt.Sprintf("query: %s", query)))

		os.Exit(constants.ExitPrepareStmtFail)
	}

	return s
}

func PrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	s, err := db.PrepareNamed(query)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("error while preparing named statement: %s", err)))
		alog.Logger.Error(errors.New(fmt.Sprintf("query: %s", query)))

		os.Exit(constants.ExitPrepareStmtFail)
	}

	return s
}
