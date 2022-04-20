package component

import "fmt"

type WheelBig struct {

}
func (t WheelBig) AddWheel() string{
	fmt.Println("Add Big Wheel")
	return "Big"
}
