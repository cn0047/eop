package controller

import (
	"fmt"
	"sync"

	"app/examples/panic_goroutine/service"
)

// SignUp - entrypoint for panic_goroutine implementation.
func SignUp(username string) {
	defer func() {
		msg := "ok"
		if r := recover(); r != nil {
			msg = r.(error).Error()
		}
		fmt.Printf("[panic_goroutine] SignUp: %s \n", msg)
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
