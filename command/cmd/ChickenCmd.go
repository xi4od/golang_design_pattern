package cmd

type ChickenCmd struct {
	Command
}

func (t ChickenCmd) ExeCommand() {
	t.Baker.BakeChicken()
}
