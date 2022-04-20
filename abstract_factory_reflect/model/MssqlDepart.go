package model

import (
	"abstract_factory_reflect/global"
	"fmt"
	"reflect"
)

type MssqlDepart struct {

}

func init() {
	v := MssqlDepart{}
	global.TypeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
}

func(t MssqlDepart)	Insert(depart Depart) {
	fmt.Printf("Depart: %v, 插入Mssql数据库\n", depart.DepartName)
}
func(t MssqlDepart) GetDepart(id int) {
	fmt.Printf("获取Depart %v by Mssql\n", id)
}