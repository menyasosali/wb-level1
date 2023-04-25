package main

import "fmt"

func main() {
	var x = 10
	var y = 50

	x, y = y, x

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v", y)

}
