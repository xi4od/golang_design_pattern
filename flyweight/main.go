package main

import (
	"flyweight/flyweight"
	"fmt"
)

func main() {
	factory := flyweight.NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xiang beauty")
	fmt.Println(len(factory.Pool))
	fmt.Println(hong)
	fmt.Println(xiang)
}
