package models

import "time"

type User struct {
	Email        string
	PasswordHash string
	CurrentToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Store(User) error
}

type SessionRepository interface {
	FindByToken(token string) (*User, error)
	New(User) (string, error)
}
