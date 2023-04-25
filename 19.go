package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	newStr := reverseString(s)
	fmt.Println(newStr)

}

// reverseString переворачивает строку, используея для этого преобразование в слайс рун
func reverseString(str string) string {
	runs := []rune(str)
	for i, j := 0, len(runs)-1; i < j; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]

	}
	return string(runs)
}
