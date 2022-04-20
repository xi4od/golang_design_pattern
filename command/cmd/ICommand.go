package cmd

type ICommand interface {
	ExeCommand()
}

type Command struct {
	Baker Baker
}

func (t Command) ExeCommand() {
}
