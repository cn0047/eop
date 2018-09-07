package service

import (
	"sync"

	c "app/common"
)

// MustSignUp - facade for bunch of functions.
func MustSignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.Recover([]error{c.ErrorUsernameBlank, c.ErrorUsernameAlreadyTaken}, func(err interface{}) {
		ch <- err.(error)
	})
	mustValidation(username)
	mustSignUpFacebook(username)
	mustSignUpTwitter(username)
	mustSignUpPinterest(username)
}

func mustValidation(username string) {
	if len(username) == 0 {
		panic(c.ErrorUsernameBlank)
	}
}

func mustSignUpFacebook(username string) {
	if username == "bond" {
		panic(c.ErrorUsernameAlreadyTaken)
	}
}

func mustSignUpTwitter(username string) {
	if username == "leiter" {
		panic(c.ErrorUsernameAlreadyTaken)
	}
}

func mustSignUpPinterest(username string) {
	if username == "q" {
		panic(c.ErrorUsernameAlreadyTaken)
	}
}
