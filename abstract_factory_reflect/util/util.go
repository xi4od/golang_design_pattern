package util

import (
	"abstract_factory_reflect/global"
	"reflect"
)

func MakeInstance(name string) interface{} {
	v := reflect.New(global.TypeRegistry[name]).Elem()
	return v.Interface()
}
