package main

//circular embedding compile-type error -
/*
type Interface1 interface {
	Interface2
	method1()
}
type Interface2 interface {
	Interface3
	method2()
}

type Interface3 interface {
	method3()
	Interface1
}


*/
