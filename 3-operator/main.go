package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 5

	//additional
	zAdd := x + y
	fmt.Println(zAdd)

	//substracts
	zSub := x - y
	fmt.Println(zSub)

	//multiplies
	zMul := x * y
	fmt.Println(zMul)

	//divides
	zDiv := x / y
	fmt.Println(zDiv)

	//modulus
	zMod := x % y
	fmt.Println(zMod)

	//relational
	i := 10
	j := 5

	fmt.Println(i == j)
	fmt.Println(i != j)
	fmt.Println(i > j)
	fmt.Println(i < j)
	fmt.Println(i >= j)
	fmt.Println(i <= j)

	//operator logic
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!b)

}
