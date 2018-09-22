// Package common contains basic functions
// which helps to work with `recover` in bit pleasant way.
package common

// Recover performs call cb function with recovered value.
func Recover(cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	cb(r)
}

// RecoverOne performs call cb function with recovered value
// in case when recovered value equals to e.
func RecoverOne(e error, cb func(v interface{})) {
	r := recover()

	if r == nil {
		return
	}

	if r == e {
		cb(r)
		return
	}

	panic(r)
}

// RecoverAny performs call cb function with recovered value
// in case when recovered value exists in slice errors.
func RecoverAny(errors []error, cb func(v interface{})) {
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
