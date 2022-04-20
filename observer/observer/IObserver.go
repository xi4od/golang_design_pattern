package observer

type IObserver interface {
	NotifyCallBack(string)
}

type Observer struct {
	Name string
}
//
//func NewObserver(name string, sub subject.Subject) *Observer {
//	return &Observer{name: name, sub: sub}
//}

