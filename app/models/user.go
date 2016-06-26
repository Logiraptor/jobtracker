package models

type User struct {
	Email        string
	PasswordHash string
	CurrentToken string
}
