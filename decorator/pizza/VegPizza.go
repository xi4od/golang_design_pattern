package pizza

import "fmt"

type VegPizza struct {

}

func (this VegPizza) GetPrice() float64  {
	fmt.Println("VegPizza: 60")
	return 60
}
