package service

import (
	"fmt"
	"sync"
)

// MustSignUp - facade for bunch of functions.
func MustSignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			ch <- fmt.Errorf(r.(string))
		}
	}()
	mustValidation(username)
	mustSignUpFacebook(username)
	mustSignUpTwitter(username)
	mustSignUpPinterest(username)
}

func mustValidation(username string) {
	if len(username) == 0 {
		panic("username cannot be blank")
	}
}

func mustSignUpFacebook(username string) {
	if username == "bond" {
		panic("username already taken")
	}
}

func mustSignUpTwitter(username string) {
	if username == "leiter" {
		panic("username already taken")
	}
}

func mustSignUpPinterest(username string) {
	if username == "q" {
		panic("username already taken")
	}
}
