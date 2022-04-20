package main

import (
	"fmt"
	"mediator/mediator"
)

func main() {
	stationManager := mediator.NewStationManger()
	passengerTrain := &mediator.PassengerTrain{
		Name:     fmt.Sprintf("passengerTrain%d", 1),
		Mediator: stationManager,
	}

	goodsTrain := &mediator.GoodsTrain{
		Name:     fmt.Sprintf("goodsTrain%d", 1),
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()

	fmt.Println("exit")
}
