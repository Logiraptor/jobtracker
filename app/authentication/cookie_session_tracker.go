package authentication

import (
	"jobtracker/app/models"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
)

type CookieSessionTracker struct {
	SessionName       string
	SessionRepository SessionRepository
	Store             sessions.Store
	Logger            *logrus.Logger
}

func NewCookieSessionTracker(sessionName string, logger *logrus.Logger, sessionStore sessions.Store, sessionRepo SessionRepository) CookieSessionTracker {
	return CookieSessionTracker{
		SessionName:       sessionName,
		SessionRepository: sessionRepo,
		Store:             sessionStore,
		Logger:            logger,
	}
}

func (c CookieSessionTracker) getToken(req *http.Request) (string, bool) {
	sess, err := c.Store.Get(req, c.SessionName)
	if err != nil {
		c.Logger.WithError(err).Errorln("Store.Get failed")
		return "", false
	}
	tokenI, ok := sess.Values["session"]
	if !ok {
		return "", false
	}
	token, ok := tokenI.(string)
	if !ok {
		return "", false
	}
	return token, true
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
	return c.Store.Save(req, rw, sess)
}

func (c CookieSessionTracker) Logout(rw http.ResponseWriter, req *http.Request) error {
	token, ok := c.getToken(req)
	if !ok {
		return nil
	}
	sess, err := c.Store.Get(req, c.SessionName)
	if err != nil {
		return err
	}
	delete(sess.Values, "session")
	sess.Save(req, rw)
	return c.SessionRepository.DeleteByToken(token)
}

func (c CookieSessionTracker) CurrentUser(req *http.Request) (*models.User, bool) {
	token, ok := c.getToken(req)
	if !ok {
		return nil, false
	}
	user, err := c.SessionRepository.FindByToken(token)
	if err != nil {
		c.Logger.WithError(err).Errorln("SessionRepository.FindByToken failed")
		return nil, false
	}
	return user, true
}
