package component

import "fmt"

type SeatCloth struct {

}

func (t SeatCloth) AddSeat() string {
	fmt.Println("Add Cloth Seat")
	return "Cloth"
}
