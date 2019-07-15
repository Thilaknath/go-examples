package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Return type is int, varibale type is declared after the variable name 
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favourite number is", rand.Intn(9))
	fmt.Printf("Now you have %g problems \n", math.Sqrt(9))
	fmt.Println(add(42, 13))
}

https://cecusddemo03.authentication.us10.hana.ondemand.com/login