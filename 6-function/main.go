package main

import "fmt"

func main() {

	fmt.Println("Call Function")

	x := 5
	y := 4

	z := add(x, y)
	fmt.Println("hasil z adalah ", z)

	var stx string = "hasilnya: "
	fmt.Println(stx, z)

	name := "firman"
	result := hello(name)
	fmt.Println(result)

	aSwap := "hello"
	bSwap := "golang"

	//get multiple value
	resultA, resultB := swap(aSwap, bSwap)
	fmt.Println(resultA, resultB)

	fmt.Println()

	//function as value
	addVal := func(x, y int) int {
		return x + y
	}

	//function as value
	helloVal := func(name string) string {
		return fmt.Sprintf("Hai %s", name)
	}

	fmt.Println(addVal(6, 7))
	fmt.Println(helloVal("hani"))

}

func add(x int, y int) int {
	return x + y
}

func hello(name string) string {
	// Sprintf for concat more one string
	return fmt.Sprintf("hello %s", name)
}

// multiple value
func swap(a string, b string) (string, string) {
	return b, a
}
