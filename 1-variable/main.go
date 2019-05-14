package main

import (
	"fmt"
)

func main() {

	var x int
	x = 10

	var y float64
	y = 5.5

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("static type x : %T\n", x)
	fmt.Printf("static type y : %T\n", y)

	z := "Dynamic"
	i := 50

	fmt.Println(z)
	fmt.Println(i)

	fmt.Printf("dynamic type z : %T\n", z)
	fmt.Printf("dynamic type i : %T\n", i)

	t := 6
	p := 3

	fmt.Println(t + p)

}
