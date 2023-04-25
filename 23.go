package main

import "fmt"

func main() {
	slai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fmt.Println(removeSimv(slai, 3))
}

func removeSimv(slai []int, i int) []int {
	slai = append(slai[:i], slai[i+1:]...)
	return slai
}
