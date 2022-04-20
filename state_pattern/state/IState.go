package state

import "fmt"

/*
已申请   applied
审批中	approving
复核中	reviewing
审批完成	approvied
*/

type IState interface {
	Handle(ctx *Context)
	StateChange(stateA IState, stateB IState)
}

type State struct {

}

func (t State) StateChange(stateA IState, stateB IState) {
	fmt.Printf("Change %T to %T\n", stateA, stateB)
}

func(t State) Handle(ctx Context){

}

