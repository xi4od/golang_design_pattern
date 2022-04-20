package state

type Context struct {
	State IState
}

func NewContext(State IState) *Context {
	return &Context{State}
}

func(t *Context) Request(){
	t.State.Handle(t)
}
