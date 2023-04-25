package main

import "fmt"

func setBit(num int64, i int64, bit int64) int64 {
	if bit == 1 {
		num |= 1 << i // установка i-го бита в 1
	} else {
		num &= ^(1 << i) // установка i-го бита в 0
	}

	return num
}

func main() {
	var num int64 = 10
	fmt.Printf("%064b\n", num)
	num = setBit(num, 0, 1)
	fmt.Printf("%064b\n", num)
	num = setBit(num, 2, 1)
	fmt.Printf("%064b\n", num)
}
