package operations

type OperationDivide struct {
}

func (this *OperationDivide) GetResult(numA float64, numB float64) float64 {
	return numA / numB
}
