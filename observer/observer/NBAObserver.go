package observer

import (
	"fmt"
)

type NBAObserver struct {
	Observer
}

func NewNBAObserver(name string) *NBAObserver {
	return &NBAObserver{Observer{name}}
}

func (t NBAObserver) NotifyCallBack(e string)  {
	fmt.Printf("%v %v 关闭NBA,继续工作\n", e, t.Name)
}
