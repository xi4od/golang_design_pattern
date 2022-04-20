package state

type ReviewingState struct {
	State
}
func(t ReviewingState) Handle(ctx *Context) {
	ctx.State = ApproviedState{}
	t.StateChange(t, ctx.State)
}
