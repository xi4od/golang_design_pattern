package cash_strategy

type CashRebate struct {
	Condition float64
	Rebate    float64
}

func (this *CashRebate) AcceptCash(price float64) float64 {
	if price >= this.Condition {
		return price - float64(int(price/this.Condition))*this.Rebate
	} else {
		return price
	}
}
