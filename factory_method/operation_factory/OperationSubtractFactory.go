package operation_factory

import "factory_method/operation"

type OperationSubtractFactory struct {
}

func (this OperationSubtractFactory) CreateOperation() operation.IOperation {
	return &operation.OperationSubtract{}
}
