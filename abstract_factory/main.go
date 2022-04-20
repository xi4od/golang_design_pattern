package main

import (
	"abstract_factory/db_factory"
	"abstract_factory/model"
)

func main(){
	user := model.User{1,"xi4od"}
	depart := model.Depart{1,"infosec"}

	var dbFactory db_factory.IFactory
	//dbFactory = db_factory.MysqlFactory{}
	dbFactory = db_factory.MssqlFactory{}

	var iUser model.IUser
	iUser = dbFactory.CreateUser()
	iUser.Insert(user)
	iUser.GetUser(1)

	var iDepart model.IDepart
	iDepart = dbFactory.CreateDepart()
	iDepart.Insert(depart)
	iDepart.GetDepart(2)
}

