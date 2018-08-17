package service

import (
	"sync"

	ec "app/panicroutinepro/taxonomy"
	"app/panicroutinepro/common"
)

func MustSignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer common.Catch([]error{ec.ErrorUsernameBlank, ec.ErrorUsernameAlreadyTaken}, func(err interface{}) {
		ch <- err.(error)
	})
	MustValidation(username)
	MustSignUpFacebook(username)
	MustSignUpTwitter(username)
	MustSignUpPinterest(username)
}

func MustValidation(username string) {
	if len(username) == 0 {
		panic(ec.ErrorUsernameBlank)
	}
}

func MustSignUpFacebook(username string) {
	if username == "bond" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}

func MustSignUpTwitter(username string) {
	if username == "leiter" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}

func MustSignUpPinterest(username string) {
	if username == "q" {
		panic(ec.ErrorUsernameAlreadyTaken)
	}
}
