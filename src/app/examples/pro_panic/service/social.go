package service

import (
	c "app/common"
)

// MustSignUp - facade for bunch of functions.
func MustSignUp(username string) {
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
