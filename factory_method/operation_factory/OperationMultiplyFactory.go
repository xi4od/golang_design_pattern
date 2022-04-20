package operation_factory

import "factory_method/operation"

type OperationMultiplyFactory struct {
}

func (this OperationMultiplyFactory) CreateOperation() operation.IOperation {
	return &operation.OperationMultiply{}
}
