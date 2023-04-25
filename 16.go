package main

import (
	"fmt"
	"sort"
)

func main() {
	mas := []int{8, 9, 10, 1, 2, 5, 6, 3, 4, 7}
	sort.Slice(mas, func(i, j int) bool {
		return mas[i] < mas[j]
	})
	fmt.Println(mas)
}
