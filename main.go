package main

import (
	"fmt"
)

var intCh chan int = make(chan int)

func sqr(number *int) {
	*number = *number + 0
	intCh <- *number
}

func sqr2() {
	var number int = <-intCh
	number = number * number
	intCh <- number
}

func main() {

	mas := [5]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i < 5; i++ {
		go sqr(&mas[i])
		go sqr2()
		sum += <-intCh
	}

	fmt.Scanln()
	fmt.Println("SUM: ", sum)
}
