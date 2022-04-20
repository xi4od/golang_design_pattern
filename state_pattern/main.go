package main

import (
	"state_pattern/state"
)

func main(){
	var mState state.IState
	mState = state.AppliedState{}
	ctx := state.NewContext(mState)
	ctx.Request()
	ctx.Request()
	ctx.Request()
	ctx.Request()
	ctx.Request()
}
