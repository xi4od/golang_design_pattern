package builder

import (
	"builder/component"
)

type Es6Builder struct {
	CarBuilder
}

func (t *Es6Builder) BuildComponent() {
	var seat component.ISeat
	var wheel component.IWheel
	seat = component.SeatCloth{}
	wheel = component.WheelSmall{}
	t.name = "ES6"
	t.seat = seat.AddSeat()
	t.wheel = wheel.AddWheel()
}