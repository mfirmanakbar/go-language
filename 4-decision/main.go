package main

import (
	"fmt"
)

func main() {

	//if else
	x := true

	if !x {
		fmt.Println("hello firman")
	} else {
		fmt.Println("hello akbar")
	}

	y := 100

	if y < 100 {
		fmt.Println("go < 100")
	} else if y > 100 {
		fmt.Println("go > 100")
	} else {
		fmt.Println("go = 100")
	}

	//switch
	a := 90

	switch a {
	case 60:
		fmt.Println("switch case 60")
	case 90:
		fmt.Println("switch case 90")
	default:
		fmt.Println("switch case not found")
	}

	switch {
	case a == 60:
		fmt.Println("switch case 60")
	case a == 90:
		fmt.Println("switch case 90")
	default:
		fmt.Println("switch case not found")
	}

	//type switch
	var b interface{}
	b = "firman"

	switch b.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case float64:
		fmt.Println("float64 type")
	default:
		fmt.Println("type not found")
	}

}
