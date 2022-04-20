package cash

import cs "strategy/cash/cash_strategy"

type CashContext struct {
	iCash ICash
}

func (this *CashContext) GetResult(price float64) float64 {
	return this.iCash.AcceptCash(price)
}

func NewCashContext(style string) *CashContext {
	context := &CashContext{}
	switch style {
	case "正常收费":
		context.iCash = &cs.CashNormal{}
	case "满300返100":
		context.iCash = &cs.CashRebate{300, 100}
	case "8折":
		context.iCash = &cs.CashDiscount{0.8}
	default:
		panic("计价方式错误！")
	}

	return context
}