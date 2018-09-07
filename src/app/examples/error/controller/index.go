package controller

import (
	"fmt"

	"app/examples/error/service"
)

// SignUp - entrypoint for error implementation.
func SignUp(username string) {
	msg := "ok"
	if err := service.SignUp(username); err != nil {
		msg = err.Error()
	}
	fmt.Printf("[error] SignUp: %s \n", msg)
}
