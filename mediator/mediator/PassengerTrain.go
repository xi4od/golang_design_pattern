package mediator

import "fmt"

// 旅客列车
type PassengerTrain struct {
	Name     string
	Mediator IMediator
}

func (g *PassengerTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println(g.Name + ": Landing")
	} else {
		fmt.Println(g.Name + ": Waiting")
	}
}

func (g *PassengerTrain) Departure() {
	fmt.Println(g.Name + ": Leaving")
	g.Mediator.NotifyFree()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println(g.Name + ": Arrival Permitted. Landing")
}
