package main

import (
	"database/sql"
	"jobtracker/app/db/migrations/util"
)

// Up is executed when this migration is applied
func Up_20160613222742(txn *sql.Tx) {
	util.Must(txn.Exec(`CREATE TABLE reset_tokens (
		user_id INT REFERENCES users (id),
		value TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
		PRIMARY KEY (user_id)
	)`))
}

// Down is executed when this migration is rolled back
func Down_20160613222742(txn *sql.Tx) {
	util.Must(txn.Exec(`DROP TABLE reset_tokens`))
}
