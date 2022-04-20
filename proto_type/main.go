package main

import (
	"fmt"
	"proto_type/proto"
)

func main(){
	aResume := proto.Resume{"xi4od","16","police","male"}
	//bResume := aResume.ClonePtr()
	bResume := aResume.Clone()
	bResume.School = "xxx"
	fmt.Printf("%v\n", aResume)
	fmt.Printf("%v\n", bResume)
}
