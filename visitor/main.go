package main

import "vistor/visitor"

func main() {
	e := new(visitor.Element)
	e.Accept(new(visitor.WeiBoVisitor))
	e.Accept(new(visitor.IQIYIVisitor))
}
