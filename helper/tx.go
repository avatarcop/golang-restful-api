package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback, "error helper/tx at func commitorrollback when rollback")
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit, "error helper/tx at func commitorrollback when commit")
	}
}
