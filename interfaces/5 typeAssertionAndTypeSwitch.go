package main

import (
	"fmt"
	"math"
)

// declaring a method for circle type
func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {

	var s Shape = Circle{radius: 2.5}

	fmt.Printf("%T\n", s)

	fmt.Printf("Circle Area:%v\n", s.Area())

	//** TYPE Assertion **//
	fmt.Printf("Sphere Volume:%v\n", s.(Circle).volume())

	ball, ok := s.(Circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	switch value := s.(type) {
	case Circle:
		fmt.Printf("%#v has circle type\n", value)
	case Rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}
