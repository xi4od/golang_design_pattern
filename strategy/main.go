package main

import (
	"fmt"
	"strategy/cash"
)

func main(){
	var total float64
	cash1 := cash.NewCashContext("正常收费")
	total += cash1.GetResult(1 * 10000)
	cash2 := cash.NewCashContext("满300返100")
	total += cash2.GetResult(1 * 10000)
	cash3 := cash.NewCashContext("8折")
	total += cash3.GetResult(1 * 10000)
	fmt.Println("total:", total)
}
