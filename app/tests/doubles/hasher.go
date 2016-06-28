package doubles

import (
	"crypto/md5"
	"encoding/hex"
)

func NewFakePasswordHasher() *FakePasswordHasher {
	return &FakePasswordHasher{
		New_: func(password string) string {
			return hex.EncodeToString(md5.New().Sum([]byte(password)))
		},
		Verify_: func(hash, password string) bool {
			return hash == hex.EncodeToString(md5.New().Sum([]byte(password)))
		},
	}
}

func (f *FakePasswordHasher) New(password string) string {
	return f.New_(password)
}

func (f *FakePasswordHasher) Verify(hash, password string) bool {
	return f.Verify_(hash, password)
}
