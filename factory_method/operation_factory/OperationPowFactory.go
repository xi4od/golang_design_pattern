package operation_factory

import "factory_method/operation"

type OperationPowFactory struct {
}

func (this OperationPowFactory) CreateOperation() operation.IOperation {
	return &operation.OperationPow{}
}
