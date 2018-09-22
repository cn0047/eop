package controller

import (
	"fmt"

	c "app/common"
	"app/examples/pro_panic/service"
)

// SignUp - entrypoint for pro_panic implementation.
func SignUp(username string) {
	defer c.Recover(func(err interface{}) {
		fmt.Printf("[pro_panic] SignUp: %s \n", err.(error))
	})
	service.MustSignUp(username)
	fmt.Printf("[pro_panic] SignUp: %s \n", "ok")
}
