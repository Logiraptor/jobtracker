package authentication

import (
	"golang.org/x/crypto/bcrypt"
)

type BCryptPasswordHasher struct {
	Cost int
}

func NewBCryptPasswordHasher(cost int) BCryptPasswordHasher {
	return BCryptPasswordHasher{Cost: cost}
}

func (b BCryptPasswordHasher) New(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.Cost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func (b BCryptPasswordHasher) Verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
