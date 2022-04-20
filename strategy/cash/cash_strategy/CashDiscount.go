package cash_strategy

type CashDiscount struct {
	Discount float64
}

func (this *CashDiscount) AcceptCash(price float64) float64 {
	if price < 0 || this.Discount < 0 {
		panic("不允许负数作为参数")
	}
	if this.Discount >= 1 {
		panic("折扣无效")
	}
	return price * this.Discount
}
