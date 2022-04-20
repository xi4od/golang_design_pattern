package pizza

import "fmt"

type MeatPizza struct {

}

func (this MeatPizza) GetPrice() float64  {
	fmt.Println("MeatPizza: 100")
	return 100
}