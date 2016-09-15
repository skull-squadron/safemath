package safemath

type Error string

func IsError(x interface{}) (yes bool) {
	_, yes = x.(Error)
	return
}

func (err Error) Error() string { return string(err) }
