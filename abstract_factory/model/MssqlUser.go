package model

import "fmt"

type MssqlUser struct {

}

func(t MssqlUser) Insert(user User) {
	fmt.Printf("User: %v, 插入Mysql数据库\n", user.Name)
}

func(t MssqlUser) GetUser(id int) {
	fmt.Printf("获取User %v by Mssql\n", id)
}
