package main

import (
	"abstract_factory_reflect/data_access"
	"abstract_factory_reflect/db_factory"
	"abstract_factory_reflect/model"
	"fmt"
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

	fmt.Println("########")
	var iDepart model.IDepart
	iDepart = dbFactory.CreateDepart()
	iDepart.Insert(depart)
	iDepart.GetDepart(2)

	fmt.Println("####reflect####")
	var iUser1 model.IUser
	da := data_access.DataAccess{}
	iUser1 = da.CreateUser()
	iUser1.Insert(user)
	iUser1.GetUser(1)

	fmt.Println("####reflect####")
	var iDepart1 model.IDepart
	iDepart1 = da.CreateDepart()
	iDepart1.Insert(depart)
	iDepart1.GetDepart(1)
}