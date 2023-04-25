package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dataChan := make(chan int)

	var workers int
	fmt.Scan(&workers)

	for i := 0; i < workers; i++ {
		go func() {
			// считывание данных из канала и вывод
			for data := range dataChan {
				fmt.Println(data)
			}
		}()
	}

	i := 0
	for {
		i++
		dataChan <- i // постоянная запись данных в канал
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	close(dataChan)
}
