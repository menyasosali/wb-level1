package main

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temps []float64) map[int][]float64 {
	result := make(map[int][]float64)

	for _, temp := range temps {
		// Вычисляем значение группы для текущей температуры
		group := int(math.Floor(temp/10)) * 10

		// Добавляем температуру в группу
		result[group] = append(result[group], temp)
	}

	// Сортируем группы по возрастанию
	var groups []int
	for group := range result {
		groups = append(groups, group)
	}
	sort.Ints(groups)

	// Возвращаем результат
	resultMap := make(map[int][]float64, len(groups))
	for _, group := range groups {
		resultMap[group] = result[group]
	}

	return resultMap
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temperatures)
	fmt.Println(groups)
}
