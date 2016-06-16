package main

import (
	"jobtracker/app/db/migrations/util"

	"database/sql"
)

// Up is executed when this migration is applied
func Up_20160613214335(txn *sql.Tx) {
	util.Must(txn.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		current_token TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
		updated_at TIMESTAMP NOT NULL
	)`))
}

// Down is executed when this migration is rolled back
func Down_20160613214335(txn *sql.Tx) {
	util.Must(txn.Exec(`DROP TABLE USERS`))
}
