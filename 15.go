package main

// Этот привести к выделению большого количества памяти в функции createHugeString,
// но затем немедленно освобождению её части, которая будет храниться в переменной justString.
// Таким образом, выделение памяти будет неоправданно,
// что может привести к проблемам с производительностью и использованием ресурсов.

var justString []byte

func someFunc() {
	v := createHugeString(1 << 10)
	justString = make([]byte, 100) // ссоздаем слайс байт с нужным размером
	copy(justString, v)            // копируем значения в justString
}

func main() {
	someFunc()
}
