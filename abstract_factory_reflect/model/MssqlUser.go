package model

import (
	"abstract_factory_reflect/global"
	"fmt"
	"reflect"
)

type MssqlUser struct {

}

func init() {
	v := MssqlUser{}
	global.TypeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
}

func(t MssqlUser) Insert(user User) {
	fmt.Printf("User: %v, 插入Mssql数据库\n", user.Name)
}

func(t MssqlUser) GetUser(id int) {
	fmt.Printf("获取User %v by Mssql\n", id)
}
