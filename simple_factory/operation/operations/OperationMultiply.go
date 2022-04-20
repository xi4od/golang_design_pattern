package operations

type OperationMultiply struct {
}

func (this *OperationMultiply) GetResult(numA float64, numB float64) float64 {
	return numA * numB
}
