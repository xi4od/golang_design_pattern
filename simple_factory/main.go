package main

import (
	"fmt"
	"simple_factory/operation"
)

func main() {
	factory := operation.NewOperationFactory()
	opt := factory.CreateOperation("+")
	fmt.Println(opt.GetResult(2,3.4))
}
