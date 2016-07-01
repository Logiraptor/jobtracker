package authentication

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"jobtracker/app/models"
	"time"
)

type PSQLUserRepo struct {
	DB *sql.DB
}

func NewPSQLUserRepo(db *sql.DB) PSQLUserRepo {
	return PSQLUserRepo{DB: db}
}

func (p PSQLUserRepo) FindByEmail(email string) (*models.User, error) {
	row := p.DB.QueryRow(`
	SELECT email, password_hash, current_token
	FROM users WHERE email = $1`, email)
	var user models.User
	err := row.Scan(&user.Email, &user.PasswordHash, &user.CurrentToken)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p PSQLUserRepo) Store(user models.User) error {
	now := time.Now()
	_, err := p.DB.Exec(`INSERT INTO users
		(email, password_hash, current_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)`,
		user.Email, user.PasswordHash, user.CurrentToken, now, now)
	return err
}

type PSQLSessionRepo struct {
	DB *sql.DB
}

func NewPSQLSessionRepo(db *sql.DB) PSQLSessionRepo {
	return PSQLSessionRepo{DB: db}
}

func (p PSQLSessionRepo) FindByToken(token string) (*models.User, error) {
	row := p.DB.QueryRow(`
	SELECT email, password_hash, current_token
	FROM users
	INNER JOIN sessions ON sessions.user_id = users.id
	WHERE sessions.token = $1
	`, token)
	var user models.User
	err := row.Scan(&user.Email, &user.PasswordHash, &user.CurrentToken)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p PSQLSessionRepo) DeleteByToken(token string) error {
	_, err := p.DB.Exec(`
	DELETE FROM sessions WHERE token = $1
	`, token)
	return err
}

// User has to exist in the database
func (p PSQLSessionRepo) New(user models.User) (string, error) {
	var buf [64]byte
	rand.Read(buf[:])

	token := hex.EncodeToString(buf[:])

	now := time.Now()
	_, err := p.DB.Exec(`
	INSERT INTO sessions (token, user_id, updated_at)
	VALUES ($1, (SELECT id FROM users WHERE email = $2), $3)`, token, user.Email, now)
	if err != nil {
		return "", err
	}
	return token, nil
}
