package state

type ApproviedState struct {
	State
}
func(t ApproviedState) Handle(ctx *Context) {
	ctx.State = ApprovingState{}
	t.StateChange(t, ctx.State)
}
