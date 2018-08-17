package service

import (
	"sync"
	"fmt"
)

func MustSignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			ch <- fmt.Errorf(r.(string))
		}
	}()
	MustValidation(username)
	MustSignUpFacebook(username)
	MustSignUpTwitter(username)
	MustSignUpPinterest(username)
}

func MustValidation(username string) {
	if len(username) == 0 {
		panic("username cannot be blank")
	}
}

func MustSignUpFacebook(username string) {
	if username == "bond" {
		panic("username already taken")
	}
}

func MustSignUpTwitter(username string) {
	if username == "leiter" {
		panic("username already taken")
	}
}

func MustSignUpPinterest(username string) {
	if username == "q" {
		panic("username already taken")
	}
}
