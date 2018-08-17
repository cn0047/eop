package service

import (
	"fmt"
	"sync"
)

func SignUp(username string, ch chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := Validation(username); err != nil {
		ch <- fmt.Errorf("validation failed, error: %s", err)
		return
	}
	if err := SignUpFacebook(username); err != nil {
		ch <- fmt.Errorf("facebook sign up failed, error: %s", err)
		return
	}
	if err := SignUpTwitter(username); err != nil {
		ch <- fmt.Errorf("twitter sign up failed, error: %s", err)
		return
	}
	if err := SignUpPinterest(username); err != nil {
		ch <- fmt.Errorf("pinterest sign up failed, error: %s", err)
		return
	}
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
