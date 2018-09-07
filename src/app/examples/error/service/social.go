package service

import (
	"fmt"
)

// SignUp - facade for bunch of functions.
func SignUp(username string) error {
	if err := validation(username); err != nil {
		return fmt.Errorf("validation failed, error: %s", err)
	}
	if err := signUpFacebook(username); err != nil {
		return fmt.Errorf("facebook sign up failed, error: %s", err)
	}
	if err := signUpTwitter(username); err != nil {
		return fmt.Errorf("twitter sign up failed, error: %s", err)
	}
	if err := signUpPinterest(username); err != nil {
		return fmt.Errorf("pinterest sign up failed, error: %s", err)
	}
	return nil
}

func validation(username string) error {
	if len(username) == 0 {
		return fmt.Errorf("username cannot be blank")
	}
	return nil
}

func signUpFacebook(username string) error {
	if username == "bond" {
		return fmt.Errorf("username already taken")
	}
	return nil
}

func signUpTwitter(username string) error {
	if username == "leiter" {
		return fmt.Errorf("username already taken")
	}
	return nil
}

func signUpPinterest(username string) error {
	if username == "q" {
		return fmt.Errorf("username already taken")
	}
	return nil
}
