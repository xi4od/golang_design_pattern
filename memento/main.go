package main

import (
	"time"
	"memento/game"
)

func main() {
	gameUser := game.GameUser{}
	mementoSaver := game.MementoSaver{}
	mementoSaver.Mementos = append(mementoSaver.Mementos, gameUser.CreateMemento())

	time.Sleep(time.Duration(2)*time.Second)

	gameUser.State = time.Now().String()
	gameUser.Show()

	gameUser.ReloadMemento(mementoSaver.Mementos[0])
	gameUser.Show()
}
