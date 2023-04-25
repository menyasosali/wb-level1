package main

import (
	"fmt"
)

// binarySearch - функция принимает на вход отсортированный по возрастанию массив целых чисел и значение целевого элемента, который нужно найти
// Возвращает позицию нужного элемента, либо -0
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(binarySearch(arr, 10))
}
