package main

import (
	"fmt"

)

func add (x int, y int) int { //creating a function and naming it 
	return x + y
}

func main () {
	fmt.Println(add(42, 13)) //calling the named function

}