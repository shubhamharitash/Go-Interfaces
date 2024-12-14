package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func Print(s Shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	c1 := Circle{radius: 5}
	r1 := Rectangle{width: 3, height: 2}

	Print(c1)
	Print(r1)

	a := c1.Area()
	fmt.Println("Circle Area:", a)
}
