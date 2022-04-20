package data_access

import (
	"abstract_factory_reflect/global"
	"abstract_factory_reflect/model"
	"abstract_factory_reflect/util"
	"fmt"
)

type DataAccess struct {

}

func(t DataAccess) CreateUser() model.IUser {
	myStruct := util.MakeInstance(fmt.Sprintf("model.%vUser", global.DBType))
	//fmt.Printf("%T\n", myStruct.(model.IUser))
	return myStruct.(model.IUser)
}

func(t DataAccess) CreateDepart() model.IDepart {
	myStruct := util.MakeInstance(fmt.Sprintf("model.%vDepart", global.DBType))
	return myStruct.(model.IDepart)
}

