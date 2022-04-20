package main

import "builder/builder"

func main() {
	var carBuilder builder.IBuilder
	carBuilder = &builder.Es6Builder{}
	carBuilder.BuildComponent()
	carBuilder.GetResult()

	var carBuilder1 builder.IBuilder
	carBuilder1 = &builder.Es8Builder{}
	carBuilder1.BuildComponent()
	carBuilder1.GetResult()
}
