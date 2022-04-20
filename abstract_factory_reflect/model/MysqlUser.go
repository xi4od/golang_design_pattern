package model

import (
	"abstract_factory_reflect/global"
	"fmt"
	"reflect"
)

type MysqlUser struct {

}

func init() {
	v := MysqlUser{}
	global.TypeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
}

func(t MysqlUser) Insert(user User) {
	fmt.Printf("User: %v, 插入Mysql数据库\n", user.Name)
}

func(t MysqlUser) GetUser(id int) {
	fmt.Printf("获取User %v by Mysql\n", id)
}
