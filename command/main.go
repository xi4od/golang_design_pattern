package main

import (
	"command/cmd"
	"command/person"
)

func main() {
	baker := cmd.Baker{}

	bakerCmd := cmd.Command{baker}

	chickenCmd1 := cmd.ChickenCmd{bakerCmd}
	sheepCmd1 := cmd.SheepCmd{bakerCmd}
	chickenCmd2 := cmd.ChickenCmd{bakerCmd}
	sheepCmd2 := cmd.SheepCmd{bakerCmd}

	waiter := person.Waiter{}
	waiter.SetOrder(chickenCmd1)
	waiter.SetOrder(sheepCmd1)
	waiter.SetOrder(chickenCmd2)
	waiter.SetOrder(sheepCmd2)

	waiter.CacelOrder(sheepCmd1)
	waiter.Notify()
}
