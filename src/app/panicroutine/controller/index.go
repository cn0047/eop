package controller

import (
	"sync"
	"fmt"

	"app/panicroutine/service"
)

func SignUp(username string) {
	defer func() {
		msg := "ok"
		if r := recover(); r != nil {
			msg = r.(error).Error()
		}
		fmt.Printf("[panicroutine] SignUp: %s \n", msg)
	}()
	performMustSignUp(username)
}

func performMustSignUp(username string) {
	ch := make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go service.MustSignUp(username, ch, wg)

	wg.Wait()
	close(ch)

	if err := <-ch; err != nil {
		panic(err)
	}
}
