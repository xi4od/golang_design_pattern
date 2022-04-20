package abstraction

import (
	"bridge/implementor"
	"fmt"
)

type BigBrushPen struct {
	implementor.IColor
}

func NewBigBrushPen(t string, color implementor.IColor) IBrushPen {
	switch t {
	case "BIG":
		return BigBrushPen{
			color,
		}
	case "SMALL":
		return SmallBrushPen{
			color,
		}
	default:
		return nil
	}
}

func (bbp BigBrushPen) DrawPicture() {
	fmt.Println("Draw picture with big brush pen")
	bbp.Use()
}

type SmallBrushPen struct {
	implementor.IColor
}

func (sbp SmallBrushPen) DrawPicture() {
	fmt.Println("Draw picture with small brush pen")
	sbp.Use()
}
