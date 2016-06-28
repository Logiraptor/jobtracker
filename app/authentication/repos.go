package authentication

import (
	"jobtracker/app/models"
)

type PSQLUserRepo struct {
}

func (p *PSQLUserRepo) FindByEmail(email string) (*models.User, error) {
	return new(models.User), nil
}

func (p *PSQLUserRepo) Store(user models.User) error {
	return nil
}
