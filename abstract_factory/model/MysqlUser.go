package model

import "fmt"

type MysqlUser struct {

}

func(t MysqlUser) Insert(user User) {
	fmt.Printf("User: %v, 插入Mysql数据库\n", user.Name)
}

func(t MysqlUser) GetUser(id int) {
	fmt.Printf("获取User %v by Mysql\n", id)
}
