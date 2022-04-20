package operation_factory

import "factory_method/operation"

type OperationAddFactory struct {
}

func (this OperationAddFactory) CreateOperation() operation.IOperation {
	return &operation.OperationAdd{}
}
