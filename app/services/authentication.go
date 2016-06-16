package services

import (
	"jobtracker/app/models"
)

//go:generate mockgen -source $GOFILE -destination ../mocks/mock_services.go -package mocks

type AuthService interface {
	Create(user models.User, password string) error
	Authenticate(email, password string) (*models.User, error)
}
