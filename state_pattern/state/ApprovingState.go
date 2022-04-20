package state

type ApprovingState struct {
	State
}
func(t ApprovingState) Handle(ctx *Context) {
	ctx.State = ReviewingState{}
	t.StateChange(t, ctx.State)
}
