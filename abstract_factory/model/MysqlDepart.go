package model

import "fmt"

type MysqlDepart struct {

}

func(t MysqlDepart)	Insert(depart Depart) {
	fmt.Printf("Depart: %v, 插入Mysql数据库\n", depart.DepartName)
}
func(t MysqlDepart) GetDepart(id int) {
	fmt.Printf("获取Depart %v by Mysql\n", id)
}
