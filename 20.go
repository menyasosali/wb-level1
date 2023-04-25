package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reversWordsInString(str))
}

// reverseString разбивает строку на слайс слов, затем переворачивает их порядок и возвращает строку слов в перевернутом порядке
func reversWordsInString(str string) string {
	mas := strings.Split(str, " ")
	for i, j := 0, len(mas)-1; i < len(mas)/2; i, j = i+1, j-1 {
		mas[i], mas[j] = mas[j], mas[i]
	}
	finalStr := strings.Join(mas, " ")
	return finalStr
}
