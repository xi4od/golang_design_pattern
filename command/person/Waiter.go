package person

import (
	"command/cmd"
	"reflect"
)

type Waiter struct {
	commands []cmd.ICommand
}

func (t *Waiter) SetOrder(command cmd.ICommand) {
	t.commands = append(t.commands, command)
}

func (t *Waiter) Notify() {
	for _, command := range t.commands {
		command.ExeCommand()
	}
}

func (t *Waiter) CacelOrder(command cmd.ICommand) {
	for i, cmd := range t.commands {
		if reflect.DeepEqual(cmd, command) {
			t.commands = append(t.commands[:i], t.commands[i+1:]...)
			break
		}
	}
}
