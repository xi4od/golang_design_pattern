package builder

import "fmt"

type IBuilder interface {
	BuildComponent()
	GetResult()
}

type CarBuilder struct {
	name string
	seat string
	wheel string
}

func(t CarBuilder) GetResult(){
	fmt.Printf("Car %v with %v wheel, %v seat\n", t.name, t.wheel, t.seat)
}
