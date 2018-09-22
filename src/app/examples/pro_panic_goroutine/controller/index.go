package controller

import (
	"fmt"
	"sync"

	"app/common"
	"app/examples/pro_panic_goroutine/service"
)

// SignUp - entrypoint for pro_panic_goroutine implementation.
func SignUp(username string) {
	defer common.Recover(func(err interface{}) {
		fmt.Printf("[pro_panic_goroutine] SignUp: %s \n", err.(error).Error())
	})
	performMustSignUp(username)
	fmt.Printf("[pro_panic_goroutine] SignUp: %s \n", "ok")
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
