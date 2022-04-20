package model

import (
	"abstract_factory_reflect/global"
	"fmt"
	"reflect"
)

type MysqlDepart struct {

}

func init() {
	v := MysqlDepart{}
	global.TypeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
}

func(t MysqlDepart)	Insert(depart Depart) {
	fmt.Printf("Depart: %v, 插入Mysql数据库\n", depart.DepartName)
}
func(t MysqlDepart) GetDepart(id int) {
	fmt.Printf("获取Depart %v by Mysql\n", id)
}
