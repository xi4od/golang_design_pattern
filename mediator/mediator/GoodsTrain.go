package mediator

import "fmt"

// 货物列车
type GoodsTrain struct {
	Name     string
	Mediator IMediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println(g.Name + ": Landing")
	} else {
		fmt.Println(g.Name + ": Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	g.Mediator.NotifyFree()
	fmt.Println(g.Name + ": Leaving")
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println(g.Name + ": Arrival Permitted. Landing")
}
