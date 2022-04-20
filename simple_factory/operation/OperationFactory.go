package operation

import  opt "simple_factory/operation/operations"

type OperationFactory struct {
}

func NewOperationFactory() *OperationFactory {
	return &OperationFactory{}
}

func (this *OperationFactory) CreateOperation(optType string) IOperation {

	switch optType {
	case "+":
		return &opt.OperationAdd{}
	case "-":
		return &opt.OperationSubtract{}
	case "*":
		return &opt.OperationMultiply{}
	case "/":
		return &opt.OperationDivide{}
	default:

	}
	return nil
}
