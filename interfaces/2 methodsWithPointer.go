package main

import "fmt"

type car struct {
	brand string
	price int
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

type distance *int

func main() {

	myCar := car{brand: "Audi", price: 40000}

	myCar.changeCar1("Opel", 21000)

	fmt.Println(myCar) // => {Audi 40000}

	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar) // -> {Seat 25000}

	var yourCar *car
	yourCar = &myCar      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Seat 25000}

	yourCar.changeCar2("VW", 30000)
	(*yourCar).changeCar2("VW", 30000)
	fmt.Println(*yourCar) // => {VW 30000}

	fmt.Println(myCar) // => {VW 30000}

}
