package authentication

import (
	"jobtracker/app/models"
	"net/http"

	"github.com/gorilla/sessions"
)

type CookieSessionTracker struct {
	SessionName       string
	SessionRepository SessionRepository
	Store             sessions.Store
}

func NewCookieSessionTracker(sessionName string, sessionStore sessions.Store, sessionRepo SessionRepository) CookieSessionTracker {
	return CookieSessionTracker{
		SessionName:       sessionName,
		SessionRepository: sessionRepo,
		Store:             sessionStore,
	}
}

func (c CookieSessionTracker) Login(rw http.ResponseWriter, req *http.Request, user models.User) error {
	sess, err := c.Store.Get(req, c.SessionName)
	if err != nil {
		return err
	}
	token, err := c.SessionRepository.New(user)
	if err != nil {
		return err
	}
	sess.Values["session"] = token
	return sess.Save(req, rw)
}

func (c CookieSessionTracker) CurrentUser(req *http.Request) (*models.User, bool) {
	sess, err := c.Store.Get(req, c.SessionName)
	if err != nil {
		return nil, false
	}
	tokenI, ok := sess.Values["session"]
	if !ok {
		return nil, false
	}
	token, ok := tokenI.(string)
	if !ok {
		return nil, false
	}
	user, err := c.SessionRepository.FindByToken(token)
	if err != nil {
		return nil, false
	}
	return user, true
}
