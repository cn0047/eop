package common

// Catch basic implementation which helps to work with `recover` in bit pleasant way.
func Catch(errors []error, cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	if len(errors) == 0 || inErrors(r, errors) {
		cb(r)
		return
	}

	panic(r)
}

func inErrors(e interface{}, errors []error) bool {
	for _, err := range errors {
		if e == err {
			return true
		}
	}

	return false
}
