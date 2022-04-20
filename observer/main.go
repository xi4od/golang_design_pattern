package main

import (
	"observer/observer"
	"observer/subject"
)

func main(){
	var boss subject.ISubject
	boss = subject.NewBoss("我胡汉三回来了1")

	var stockOb observer.IObserver
	stockOb = observer.NewStockObserver("xi4od")

	var NBAOb observer.IObserver
	NBAOb = observer.NewNBAObserver("kunkkun")

	boss.Attach(stockOb)
	boss.Attach(NBAOb)
	boss.Notify()

	boss.Detach(stockOb)
	boss.Notify()
}
