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
		notpassword := fake.Characters(20)
		hashed1 := hasher.New(password)
		hashed2 := hasher.New(password)
		assert.NotEqual(t, password, hashed1, "the hash should not be the password")
		assert.NotEqual(t, password, hashed2, "the hash should not be the password")
		assert.NotEqual(t, hashed1, hashed2, "multiple calls to New should yield different results")
		assert.True(t, hasher.Verify(hashed1, password), "hashes should be tied to the password that generated them")
		assert.True(t, hasher.Verify(hashed2, password), "hashes should be tied to the password that generated them")
		assert.False(t, hasher.Verify(hashed1, notpassword), "hashes should not be tied to passwords that did not generate them")
		assert.False(t, hasher.Verify(hashed2, notpassword), "hashes should not be tied to passwords that did not generate them")
	}
}
