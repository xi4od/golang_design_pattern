package component

import "fmt"

type WheelSmall struct {

}
func (t WheelSmall) AddWheel() string{
	fmt.Println("Add Small Wheel")
	return "Small"
}
