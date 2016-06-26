package doubles

func NewFakePasswordHasher() *FakePasswordHasher {
	return &FakePasswordHasher{
		New_: func(password string) string {
			return password + "hash"
		},
		Verify_: func(hash, password string) bool {
			return password+"hash" == hash
		},
	}
}

func (f *FakePasswordHasher) New(password string) string {
	return f.New_(password)
}

func (f *FakePasswordHasher) Verify(hash, password string) bool {
	return f.Verify_(hash, password)
}
