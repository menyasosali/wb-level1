package main

import "fmt"

// sovp - все значения из первого слайса записываем в мапу со значением true, проходимся по значением из 2 слайса, если
// они в мапе и значение есть, то добавляет в слайс result.
// Вычисляем l, чтобы избежать дополнительных расходов на перераспределение памяти
func sovp(arr1 []int, arr2 []int) []int {
	var l int
	if len(arr1) <= len(arr2) {
		l = len(arr2)
	} else {
		l = len(arr1)
	}
	karta := make(map[int]bool, l)

	for _, val := range arr1 {
		karta[val] = true
	}

	result := make([]int, 0, l)

	for _, val := range arr2 {
		if karta[val] == true {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	arr1 := []int{6, 1, 2, 8, 7, 9}
	arr2 := []int{1, 2, 4, 3, 5, 9}
	fmt.Println(sovp(arr1, arr2))

}
