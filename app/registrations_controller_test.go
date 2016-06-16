package app

import (
	"jobtracker/app/mocks"
	"jobtracker/app/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/manveru/faker"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestRegistrationsController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	pather := mocks.NewMockPather(mockCtrl)
	pather.EXPECT().Path("index").Return("fake path")

	fake, _ := faker.New("en")
	email := fake.Email()
	password := fake.Characters(10)

	authService := mocks.NewMockAuthService(mockCtrl)
	authService.EXPECT().Create(models.User{
		Email: email,
	}, password)

	var controller = RegistrationsController{
		Pather:      pather,
		AuthService: authService,
	}

	var recorder = httptest.NewRecorder()
	var request = mustNewRequest(t, "POST", "/create", url.Values{
		"email":            {email},
		"password":         {password},
		"password_confirm": {password},
	})

	controller.Create(recorder, request)

	assert.Equal(t, http.StatusFound, recorder.Code)
	assert.Equal(t, "/fake path", recorder.HeaderMap.Get("Location"))
}
