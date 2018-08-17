package controller

import (
	"fmt"

	"app/panic/service"
)

func SignUp(username string) {
	defer func() {
		msg := "ok"
		if r := recover(); r != nil {
			msg = r.(string)
		}
		fmt.Printf("[panic] SignUp: %s \n", msg)
	}()
	service.MustSignUp(username)
}
