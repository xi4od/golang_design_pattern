package main

import (
	"proxy/log"
	"proxy/user"
)

func main(){
	var user2 user.IUser
	user1 := user.User{Name: "xi4od"}
	user2 = log.LoginLog{user1}

	user1.Login()
	user2.Login()
}
