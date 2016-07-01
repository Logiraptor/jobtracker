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
	DeleteByToken(token string) error
	FindByToken(token string) (*models.User, error)
	New(models.User) (string, error)
}

type HTTPSessionTracker interface {
	Login(http.ResponseWriter, *http.Request, models.User) error
	Logout(http.ResponseWriter, *http.Request) error
	CurrentUser(*http.Request) (*models.User, bool)
}
