package component

import "fmt"

type SeatLeather struct {

}

func (t SeatLeather) AddSeat() string {
	fmt.Println("Add Leather Seat")
	return "Leather"
}

