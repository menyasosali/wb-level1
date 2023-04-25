package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// создание множества
	karta := make(map[string]bool, len(arr))

	// добавление элементов в множество
	for _, val := range arr {
		karta[val] = true
	}

	for key := range karta {
		fmt.Println(key)
	}
}
