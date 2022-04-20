package model

import "fmt"

type MssqlDepart struct {

}

func(t MssqlDepart)	Insert(depart Depart) {
	fmt.Printf("Depart: %v, 插入Mssql数据库\n", depart.DepartName)
}
func(t MssqlDepart) GetDepart(id int) {
	fmt.Printf("获取Depart %v by Mssql\n", id)
}