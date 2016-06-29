package doubles

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func NewFakePasswordHasher() *FakePasswordHasher {
	return &FakePasswordHasher{
		New_: func(password string) string {
			var salt [16]byte
			rand.Read(salt[:])
			return hex.EncodeToString(salt[:]) + "$" + password
		},
		Verify_: func(hash, password string) bool {
			salt := strings.Split(hash, "$")[0]
			return hash == salt+"$"+password
		},
	}
}

func (f *FakePasswordHasher) New(password string) string {
	return f.New_(password)
}

func (f *FakePasswordHasher) Verify(hash, password string) bool {
	return f.Verify_(hash, password)
}
