package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	sum := 0

	var wg sync.WaitGroup // создаем объект sync.WaitGroup для ожидания завершения всех горутин
	for _, num := range nums {
		wg.Add(1) // блокируем одну горутину в ожидании
		go func(num int) {
			defer wg.Done()  // освобождаем одну горутину из ожидания
			sum += num * num // добавляем к сумме квадрат числа
		}(num)
	}
	wg.Wait() // ждем пока значение заблокированных горутин не станет равным 0

	fmt.Println(sum)
}
