package main

import "fmt"

// generator - передаем значения из nums в канал out
func generator(nums []int, out chan<- int) {
	for i := range nums {
		out <- nums[i]
	}
	close(out)
}

// ymnoj - получаем значения из канала in, умноем каждое значение и передаем в канал out
func ymnoj(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)

	nums := []int{1, 2, 3, 4, 5, 6}

	go generator(nums, firstChan)

	go ymnoj(firstChan, secondChan)

	for i := range secondChan {
		fmt.Println(i)
	}
}
