package log

import (
	"fmt"
	"proxy/user"
)

type LoginLog struct {
	user.User
}

func (this LoginLog) Login(){

	fmt.Printf("Before %v Login!\n", this.User.Name)
	this.User.Login()
	fmt.Printf("After %v Login!\n", this.User.Name)
}
