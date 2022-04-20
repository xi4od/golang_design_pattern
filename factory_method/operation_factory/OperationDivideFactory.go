package operation_factory

import "factory_method/operation"

type OperationDivideFactory struct {
}

func (this OperationDivideFactory) CreateOperation() operation.IOperation {
	return &operation.OperationDivide{}
}
