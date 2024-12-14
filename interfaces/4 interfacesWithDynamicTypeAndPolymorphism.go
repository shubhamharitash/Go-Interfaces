package main

import (
	"fmt"
)

func main() {

	var s Shape
	fmt.Printf("%T\n", s)
	ball := Circle{radius: 2.5}
	s = ball
	Print(s)
	fmt.Printf("Type of s: %T\n", s)
	room := Rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s)
}
