package model

type IUser interface {
	Insert(user User)
	GetUser(id int)
}
