package service

import (
	"sync"

	"app/panicroutinepro/common"
	ec "app/panicroutinepro/taxonomy"
)

// MustSignUp - facade for bunch of functions.
func MustSignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer common.Recover([]error{ec.ErrorUsernameBlank, ec.ErrorUsernameAlreadyTaken}, func(err interface{}) {
		ch <- err.(error)
	})
	mustValidation(username)
	mustSignUpFacebook(username)
	mustSignUpTwitter(username)
	mustSignUpPinterest(username)
}

func mustValidation(username string) {
	if len(username) == 0 {
		panic(ec.ErrorUsernameBlank)
	}
}

func mustSignUpFacebook(username string) {
	if username == "bond" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}

func mustSignUpTwitter(username string) {
	if username == "leiter" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}

func mustSignUpPinterest(username string) {
	if username == "q" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}
