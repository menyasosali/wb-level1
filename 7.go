package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

// writeToMap - функция для конкурентной записив в мапу. Для этого мы блокируем мьютекс, записываем данные в map и разблокируем мьютекс
func writeToMap(key string, val int, karta map[string]int) {
	mu.Lock()
	defer mu.Unlock()
	karta[key] = val
}

func main() {
	karta := make(map[string]int)
	str := "a"
	for i := 1; i <= 10; i++ {
		k := i
		str += "a"
		go writeToMap(str, k, karta)
	}

	fmt.Println(karta)
}
