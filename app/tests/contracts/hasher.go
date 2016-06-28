package contracts

import (
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"

	"testing"
)

type Hasher interface {
	New(string) string
	Verify(string, string) bool
}

func PasswordHasher(t *testing.T, hasher Hasher) {
	fake, _ := faker.New("en")
	for i := 0; i < 10; i++ {
		password := fake.Characters(20)
		hashed := hasher.New(password)
		assert.NotEqual(t, password, hashed)
		assert.NotContains(t, hashed, password)
		assert.NotContains(t, password, hashed)
		assert.True(t, hasher.Verify(hashed, password))
	}
}
