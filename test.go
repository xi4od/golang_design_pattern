package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	A int
	B int
}

var typeRegistry = make(map[string]reflect.Type)

func init() {
	myTypes := []interface{}{MyStruct{}}
	for _, v := range myTypes {
		// typeRegistry["MyString"] = reflect.TypeOf(MyString{})
		typeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	}
}

func makeInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	return v.Interface()
}

func main() {
	//fmt.Println(typeRegistry)
	myStruct := makeInstance("main.MyStruct")
	fmt.Printf("%T",myStruct)
}
