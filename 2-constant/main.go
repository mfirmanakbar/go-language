package main

import (
	"fmt"
)

const A string = "Hello, this is constant variable ..."

func main() {

	fmt.Println(A)

	const X int = 10

	fmt.Println(X)

	z := 50 / X

	fmt.Println(z)

}
