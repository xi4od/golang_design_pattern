package operation

import "math"

type OperationPow struct {
}

func (this OperationPow) GetResult(numA float64, numB float64) float64 {
	return math.Pow(numA, numB)
}
