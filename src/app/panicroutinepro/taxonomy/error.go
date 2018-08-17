package taxonomy

import "fmt"

var (
	ErrorUsernameBlank = fmt.Errorf("username blank")
	ErrorUsernameAlreadyTaken = fmt.Errorf("username already taken")
)
