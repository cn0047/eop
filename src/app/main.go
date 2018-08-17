package main

import (
	"flag"

	errorController "app/error/controller"
	errorroutineController "app/errorroutine/controller"
	panicController "app/panic/controller"
	panicroutineController "app/panicroutine/controller"
	panicroutineproController "app/panicroutinepro/controller"
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
	case "errorroutine":
		errorroutineController.SignUp(username)
	case "panic":
		panicController.SignUp(username)
	case "panicroutine":
		panicroutineController.SignUp(username)
	case "panicroutinepro":
		panicroutineproController.SignUp(username)
	}
}
