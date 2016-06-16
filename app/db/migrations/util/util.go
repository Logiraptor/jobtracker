package util

func Must(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}
