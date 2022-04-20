package db

import "fmt"

type Mysql struct {

}

func(t Mysql) SaveData(){
	fmt.Println("Save data to Mysql")
}
