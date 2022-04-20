package operation

type OperationSubtract struct {
}

func (this OperationSubtract) GetResult(numA float64, numB float64) float64 {
	return numA - numB
}
