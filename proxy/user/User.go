package user

import "fmt"

type User struct {
	Name string
}

func (this *User) Login(){
	fmt.Printf("User %v Login!\n", this.Name)
}