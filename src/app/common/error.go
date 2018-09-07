package common

import "fmt"

var (
	// ErrorUsernameBlank - custom error for "blank username" cases.
	ErrorUsernameBlank        = fmt.Errorf("username blank")

	// ErrorUsernameAlreadyTaken - custom error for "username already taken" cases.
	ErrorUsernameAlreadyTaken = fmt.Errorf("username already taken")
)
