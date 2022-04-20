package cmd

type SheepCmd struct {
	Command
}

func (t SheepCmd) ExeCommand() {
	t.Baker.BakeSheep()
}
