package authentication

import (
	"jobtracker/app/tests/contracts"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptPasswordHasher(t *testing.T) {
	hasher := BCryptPasswordHasher{Cost: bcrypt.MinCost}
	contracts.PasswordHasher(t, hasher)
}
