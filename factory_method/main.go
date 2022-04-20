package main

import (
	"factory_method/operation_factory"
	"fmt"
)

func compute(factory operation_factory.IOperationFactory, a float64, b float64) float64 {
	return factory.CreateOperation().GetResult(a, b)
}

func main(){
	//oper := operation_factory.OperationAddFactory{}
	//oper := operation_factory.OperationDivideFactory{}
	//oper := operation_factory.OperationSubtractFactory{}
	//oper := operation_factory.OperationMultiplyFactory{}
	var oper operation_factory.IOperationFactory
	oper = operation_factory.OperationPowFactory{}
	fmt.Println(compute(oper, 4, 3))

}
