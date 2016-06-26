package authentication

import "jobtracker/app/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Store(models.User) error
}

type SessionRepository interface {
	FindByToken(token string) (*models.User, error)
	New(models.User) (string, error)
}
