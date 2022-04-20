package topping

import (
	"decorator/pizza"
	"fmt"
)

type TomatoTopping struct {
	Pizza pizza.IPizza
}

func (this TomatoTopping) GetPrice1() float64{
	fmt.Println("TomatoTopping: 10")
	return this.Pizza.GetPrice() + 10
}
