package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty) // => 5

	empty = "Go"
	fmt.Println(empty) // => Go

	empty = []int{2, 34, 4}
	fmt.Println(empty) // => [2 34 4]

	fmt.Println(len(empty.([]int))) // => 3

	you := person{}

	you.info = "You name"
	you.info = 40
	you.info = []float64{4.5, 6., 8.1}

	fmt.Println(you.info)
}
