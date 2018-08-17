package service

import (
	"fmt"
)

func SignUp(username string) error {
	if err := Validation(username); err != nil {
		return fmt.Errorf("validation failed, error: %s", err)
	}
	if err := SignUpFacebook(username); err != nil {
		return fmt.Errorf("facebook sign up failed, error: %s", err)
	}
	if err := SignUpTwitter(username); err != nil {
		return fmt.Errorf("twitter sign up failed, error: %s", err)
	}
	if err := SignUpPinterest(username); err != nil {
		return fmt.Errorf("pinterest sign up failed, error: %s", err)
	}
	return nil
}

func Validation(username string) error {
	if len(username) == 0 {
		return fmt.Errorf("username cannot be blank")
	}
	return nil
}

func SignUpFacebook(username string) error {
	if username == "bond" {
		return fmt.Errorf("username already taken")
	}
	return nil
}

func SignUpTwitter(username string) error {
	if username == "leiter" {
		return fmt.Errorf("username already taken")
	}
	return nil
}

func SignUpPinterest(username string) error {
	if username == "q" {
		return fmt.Errorf("username already taken")
	}
	return nil
}
