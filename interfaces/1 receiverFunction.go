package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() {

	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {

	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()

	fmt.Printf("%T\n", seconds)
	fmt.Println("Seconds in a day:", seconds)

	friends := names{"Dan", "Marry"}
	friends.print()

	names.print(friends)
}
