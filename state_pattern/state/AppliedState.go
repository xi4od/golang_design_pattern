package state

type AppliedState struct {
	State
}
func(t AppliedState) Handle(ctx *Context) {
	ctx.State = ApprovingState{}
	t.StateChange(t, ctx.State)
}
