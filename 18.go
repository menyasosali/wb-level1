package main

import (
	"fmt"
	"sync"
)

// Facture - используем мьютексы для обеспечения безопасности при конкурентном доступе
type Facture struct {
	value int
	mu    sync.Mutex
}

// Increment - инкрементирует значение счетчика
func (f *Facture) Increment() {
	f.mu.Lock()
	f.value++
	f.mu.Unlock()
}

// GetValue - возвращает значение счетчика
func (f *Facture) GetValue() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.value
}

func main() {
	f := &Facture{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			f.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(f.GetValue())

}
