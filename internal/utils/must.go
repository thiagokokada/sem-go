package utils

func Must1[T any](v T, err error) T {
	Must(err)
	return v
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
