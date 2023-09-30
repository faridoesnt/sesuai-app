package asql

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// ReleaseTx clean db transaction by commit if no error, or rollback if an error occurred
func ReleaseTx(tx *sqlx.Tx, err *error) {
	if *err != nil {
		// if an error occurred, rollback transaction
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Printf("unable to rollback transaction: %s", errRollback)
		} else {
			log.Print("transaction rolled back")
		}
		return
	}
	// else, commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		log.Printf("unable to commit transaction: %s", errCommit)
	}
}
