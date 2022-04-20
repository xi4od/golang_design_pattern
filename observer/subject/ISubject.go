package subject

import "observer/observer"

type ISubject interface {
	Attach(observer.IObserver)
	Detach(observer.IObserver)
	Notify()
}