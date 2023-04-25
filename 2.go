package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	squares := make([]int, len(nums))

	var wg sync.WaitGroup // создаем объект sync.WaitGroup для ожидания завершения всех горутин
	for i, num := range nums {
		wg.Add(1) // блокируем одну горутину в ожидании
		go func(i, num int) {
			defer wg.Done() // освобождаем одну горутину из ожидания
			squares[i] = num * num
		}(i, num)
	}
	wg.Wait() // ждем пока значение заблокированных горутин не станет равным 0

	fmt.Println(squares)
}
