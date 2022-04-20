package main

import (
	"decorator/pizza"
	"decorator/topping"
	"fmt"
)

func main() {
	var aMeatPizza pizza.IPizza
	aMeatPizza = pizza.MeatPizza{}
	aMeatPizzaWithCheese := topping.CheeseTopping{aMeatPizza}
	aMeatPizzaWithCheeseAndTomato := topping.TomatoTopping{aMeatPizzaWithCheese}
	fmt.Println(aMeatPizzaWithCheeseAndTomato.GetPrice1())
}
