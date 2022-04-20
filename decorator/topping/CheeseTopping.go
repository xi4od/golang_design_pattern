package topping

import (
	"decorator/pizza"
	"fmt"
)

type CheeseTopping struct {
	Pizza pizza.IPizza
}

func (this CheeseTopping) GetPrice() float64{
	fmt.Println("CheeseTopping: 15")
	return this.Pizza.GetPrice() + 15
}
