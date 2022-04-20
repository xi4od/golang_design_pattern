package cash_strategy

type CashNormal struct {

}

func (this *CashNormal) AcceptCash(price float64) float64 {
	if price < 0 {
		panic("不允许负数作为参数")
	}
	return price
}
