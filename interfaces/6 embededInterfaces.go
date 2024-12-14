package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type object interface {
	volume() float64
}

// geometry is embedding shape and object interfaces
type geometry interface {
	shape
	object
	getColor() string
}
type cube struct {
	edge  float64
	color string //newly added
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

// newly added func
func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	c := cube{edge: 2}

	//before adding color in cube , we get error
	//a, v:=measure (c)

	//after adding no eror
	a, v := measure(c)

	fmt.Printf("Area: %v, Volume: v\n", a, v)
}
