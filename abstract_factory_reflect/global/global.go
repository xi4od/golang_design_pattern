package global

import "reflect"

var TypeRegistry = make(map[string]reflect.Type)

var DBType = "Mssql"
