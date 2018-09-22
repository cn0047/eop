package main

import (
	"flag"

	errorController "app/examples/error/controller"
	errorGoroutineController "app/examples/error_goroutine/controller"
	panicController "app/examples/panic/controller"
	panicGoroutineController "app/examples/panic_goroutine/controller"
	proPanicController "app/examples/pro_panic/controller"
	proPanicGoroutineController "app/examples/pro_panic_goroutine/controller"
)

func main() {
	strategy := ""
	flag.StringVar(&strategy, "strategy", "error", "code execution strategy")

	username := ""
	flag.StringVar(&username, "username", "", "username")

	flag.Parse()

	switch strategy {
	case "error":
		errorController.SignUp(username)
	case "error_goroutine":
		errorGoroutineController.SignUp(username)
	case "panic":
		panicController.SignUp(username)
	case "panic_goroutine":
		panicGoroutineController.SignUp(username)
	case "pro_panic":
		proPanicController.SignUp(username)
	case "pro_panic_goroutine":
		proPanicGoroutineController.SignUp(username)
	}
}
