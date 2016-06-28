package authentication

import (
	"database/sql"
	"jobtracker/app/tests/contracts"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"testing"
)

func testDatabase(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "postgresql://app:@localhost:5433/jobtracker-test")
	assert.NoError(t, err)
	assert.NoError(t, db.Ping())
	return db
}

func TestPSQLUserRepo(t *testing.T) {
	db := testDatabase(t)
	defer db.Close()

	repo := &PSQLUserRepo{}
	contracts.UserRepo(t, repo)
}
