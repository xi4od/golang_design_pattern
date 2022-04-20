package builder

import (
	"builder/component"
)

type Es8Builder struct {
	CarBuilder
}

func (t *Es8Builder) BuildComponent() {
	var seat component.ISeat
	var wheel component.IWheel
	seat = component.SeatLeather{}
	wheel = component.WheelBig{}
	t.name = "ES8"
	t.seat = seat.AddSeat()
	t.wheel = wheel.AddWheel()
}
