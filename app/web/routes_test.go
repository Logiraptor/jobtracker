package web

import (
	"jobtracker/app/tests/doubles"
	"testing"

	"github.com/manveru/faker"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPather(t *testing.T) {
	var fake, _ = faker.New("en")
	var definedPathName = fake.Characters(10)
	var undefinedPathName = fake.Characters(10)

	routes := mux.NewRouter()
	routes.NewRoute().Path("/test_path").Name(definedPathName)
	var logger = &doubles.FakeLogger{}
	p := NewPather(logger, routes)
	assert.Equal(t, p.Path(definedPathName), "/test_path")
	assert.Equal(t, p.Path(undefinedPathName), "/")
	assert.Len(t, logger.Logs, 1)
	assert.Contains(t, logger.Logs[0], undefinedPathName)
}
