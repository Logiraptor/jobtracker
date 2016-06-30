package authentication

import (
	"jobtracker/app/models"
	"net/http"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Store(models.User) error
}

type SessionRepository interface {
	FindByToken(token string) (*models.User, error)
	New(models.User) (string, error)
}

type HTTPSessionTracker interface {
	Login(http.ResponseWriter, *http.Request, models.User) error
	CurrentUser(*http.Request) (*models.User, bool)
}
