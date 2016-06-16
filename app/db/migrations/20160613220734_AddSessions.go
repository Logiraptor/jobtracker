package main

import (
	"database/sql"
	"jobtracker/app/db/migrations/util"
)

// Up is executed when this migration is applied
func Up_20160613220734(txn *sql.Tx) {
	util.Must(txn.Exec(`CREATE TABLE sessions (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users (id),
		token TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
		updated_at TIMESTAMP NOT NULL
	)`))
}

// Down is executed when this migration is rolled back
func Down_20160613220734(txn *sql.Tx) {
	util.Must(txn.Exec(`DROP TABLE sessions`))
}
