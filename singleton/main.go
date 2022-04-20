package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	s3 := &singleton.Singleton{}
	fmt.Println(s1 == s2)
	fmt.Println(s2 == s3)
}
