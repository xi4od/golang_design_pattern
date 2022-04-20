package main

import (
	"bridge/abstraction"
	"bridge/implementor"
)

func main() {
	var color implementor.IColor
	color = implementor.Yellow{}
	bp := abstraction.NewBigBrushPen("BIG", color)
	bp.DrawPicture()
}
