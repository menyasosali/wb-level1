package main

import "fmt"

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(ynicalSimv(str1))
	fmt.Println(ynicalSimv(str2))
	fmt.Println(ynicalSimv(str3))
}

// ynicalSimv добавляем символ в мапу и сколько раз она встречается, если более одного раза, то воврщаем false
func ynicalSimv(str string) bool {
	slovar := make(map[rune]int)
	runs := []rune(str)
	for _, run := range runs {
		slovar[run] += 1
		if slovar[run] > 1 {
			return false
		}
	}
	return true
}
