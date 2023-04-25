package main

import "fmt"

func main() {
	var i int = 10
	var s string = "stroka teksta"
	var b bool = false
	var c chan interface{}

	getType(i)
	getType(s)
	getType(b)
	getType(c)

}

// getType определяет тип переменной и выводит его
func getType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	case chan interface{}:
		fmt.Println("Chan")
	default:
		fmt.Println("Unknown")
	}
}
