package game

import (
	"fmt"
	"time"
)

type GameUser struct {
	State string
}

func(t *GameUser) ReloadMemento(m Memento){
	t.State = m.State
}

func(t *GameUser) CreateMemento() Memento {
	t.State = time.Now().String()
	return Memento{t.State}
}

func(t *GameUser) Show() {
	fmt.Println("加载游戏存档: " + t.State)
}
