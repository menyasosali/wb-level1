package main

import "fmt"

type Human struct {
	Age    int
	Name   string
	Sex    string
	Height float64
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) GetSex() string {
	return h.Sex
}

func (h *Human) GetHeight() float64 {
	return h.Height
}

// встраивание структуры Human в структуру Action
type Action struct {
	Human
	Hobby string
}

func main() {
	h := Human{Name: "John", Sex: "Male", Age: 23, Height: 175}
	fmt.Println(h.GetName())

	a := Action{Human: h, Hobby: "Webcam"}
	// использование методов структуры Human
	fmt.Println(a.GetAge())
	fmt.Println(a.GetHeight())
	fmt.Println(a.GetSex())
	fmt.Println(a.GetName())
}
