package singleton

import "sync"

type Singleton struct {
}

var ins *Singleton
var mu sync.Mutex

func GetInstance() *Singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &Singleton{}
		}
	}
	return ins
}
