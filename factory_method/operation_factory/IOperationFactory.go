package operation_factory

import "factory_method/operation"

type IOperationFactory interface {
	CreateOperation() operation.IOperation
}
