package main

import (
	"fmt"
	"time"
)

func main() {
	digChan := make(chan int)

	var n int
	fmt.Scan(&n)

	// Горутина для записи значений в канал
	go func() {
		for i := 0; i < n; i++ {
			k := i
			digChan <- k
			time.Sleep(1 * time.Second)
		}
		close(digChan)
	}()

	// Чтение значений из канала
	for {
		select {
		case dig, ok := <-digChan:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println(dig)
		case <-time.After(5 * time.Second):
			fmt.Println("Время вышло")
			return
		}
	}
}
