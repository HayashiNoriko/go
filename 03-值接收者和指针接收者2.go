package main

import (
	"fmt"
)

func main() {
	s := Sheep{name: "喜羊羊"}
	s.Bark()
	s.Rename("懒羊羊")
	s.Bark()
}

type Sheep2 struct {
	name string
}
