package observer

import (
	"fmt"
)

type StockObserver struct {
	Observer
}

func NewStockObserver(name string) *StockObserver {
	return &StockObserver{Observer{name}}
}

func (t StockObserver) NotifyCallBack(e string)  {
	fmt.Printf("%v %v 关闭股票,继续工作\n", e, t.Name)
}
