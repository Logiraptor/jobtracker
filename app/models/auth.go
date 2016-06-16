package models

import "time"

//go:generate mockgen -source $GOFILE -destination ../mocks/mock_models.go -package mocks -imports .=jobtracker/app/models

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
