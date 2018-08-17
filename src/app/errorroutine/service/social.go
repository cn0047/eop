package service

import (
	"fmt"
	"sync"
)

// SignUp - facade for bunch of functions.
func SignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := validation(username); err != nil {
		ch <- fmt.Errorf("validation failed, error: %s", err)
		return
	}
	if err := signUpFacebook(username); err != nil {
		ch <- fmt.Errorf("facebook sign up failed, error: %s", err)
		return
	}
	if err := signUpTwitter(username); err != nil {
		ch <- fmt.Errorf("twitter sign up failed, error: %s", err)
		return
	}
	if err := signUpPinterest(username); err != nil {
		ch <- fmt.Errorf("pinterest sign up failed, error: %s", err)
		return
	}
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
