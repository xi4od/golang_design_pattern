package cash

type ICash interface {
	AcceptCash(price float64) float64
}
