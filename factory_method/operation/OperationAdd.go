package operation

type OperationAdd struct {
}

func (this OperationAdd) GetResult(numA float64, numB float64) float64 {
	return numA + numB
}
