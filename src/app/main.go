package main

import (
	"flag"

	errorController "app/examples/error/controller"
	errorGoroutineController "app/examples/error_goroutine/controller"
	panicController "app/examples/panic/controller"
	panicGoroutineController "app/examples/panic_goroutine/controller"
	panicGoroutineProController "app/examples/panic_goroutine_pro/controller"
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
	case "panic_goroutine_pro":
		panicGoroutineProController.SignUp(username)
	}
}
