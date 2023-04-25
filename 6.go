package main

import (
	"context"
	"fmt"
	"time"
)

// stopWithReturn - горутина останавливается при выходе из функции
func stopWithReturn() {
	for i := 0; i < 100; i++ {
		if i == 5 {
			fmt.Println("Значение достигнуто")
			fmt.Println("Конец stopWithReturn")
			return
		}
		fmt.Println(i)
	}
}

// stopWithReturn - горутина останавливается при отмене контекста
func stopWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Конец stopWithContext")
			return
		default:
			fmt.Println("Выполняется stopWithContext")
		}
	}
}

// stopWithChannel - горутина останавливается при передаче в канал значения
func stopWithChannel(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Конец stopWithChannel")
			return
		default:
			fmt.Println("Выполняется stopWithChannel")
		}
	}
}

// stopWithTimer - горутина останавливается после окончания времени таймера
func stopWithTimer(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("Конец stopWithTimer")
			return
		default:
			fmt.Println("Выполняется stopWithTimer")
		}
	}
}

// stopWithFlag - горутина останавливается после того, как переменная-флаг примет определенное значение
func stopWithFlag(done *bool) {
	for {
		if *done {
			fmt.Println("Конец stopWithFlag")
			return
		}
		fmt.Println("Выполняется stopWithFlag")
	}
}

func main() {
	go stopWithReturn()

	ctx, cancel := context.WithCancel(context.Background())
	go stopWithContext(ctx)
	time.Sleep(2 * time.Millisecond)
	cancel()

	stop := make(chan bool)
	go stopWithChannel(stop)
	time.Sleep(2 * time.Millisecond)
	stop <- false

	timer := time.NewTimer(1 * time.Millisecond)
	go stopWithTimer(timer)
	time.Sleep(2 * time.Millisecond)

	done := false
	go stopWithFlag(&done)
	time.Sleep(2 * time.Millisecond)
	done = true

	time.Sleep(1 * time.Second)
}
