package subject

import (
	"observer/observer"
)

type Boss struct {
	//Observers sync.Map
	Observers map[observer.IObserver] interface{}
	Action string
}

func NewBoss(action string) *Boss {
	return &Boss{Observers: make(map[observer.IObserver] interface{}), Action: action}
}

func(t *Boss) Attach(ob observer.IObserver) {
	//t.Observers.Store(ob, struct{}{})
	t.Observers[ob] = 1
}

func(t *Boss) Detach(ob observer.IObserver) {
	//t.Observers.Delete(ob)
	delete(t.Observers, ob)
}

func(t *Boss) Notify() {
	//t.Observers.Range(func(key, value interface{}) bool {
	//	if key == nil {
	//		return false
	//	}
	//	key.(observer.IObserver).NotifyCallBack(t.Action)
	//	return true
	//})

	for ob:= range t.Observers {
		ob.NotifyCallBack(t.Action)
	}
}
