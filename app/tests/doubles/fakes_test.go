package doubles

import (
	"jobtracker/app/tests/contracts"

	"testing"
)

func TestFakes(t *testing.T) {
	contracts.PasswordHasher(t, NewFakePasswordHasher())
	contracts.UserRepo(t, NewFakeUserRepository())
	contracts.SessionRepo(t, NewFakeSessionRepository())
}
