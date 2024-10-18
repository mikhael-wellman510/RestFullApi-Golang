package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	// Membuat data tetap berjalan
	// Jika tidak terjadi Panic , maka recover hasil nya <nil>
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfErr(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfErr(errorCommit)
	}
}
