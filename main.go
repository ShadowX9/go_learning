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
	fmt.Println(number, " ")
}

func main() {

	mas := [5]int{2, 4, 6, 8, 10}

	for i := 0; i < 5; i++ {
		go sqr(&mas[i])
		go sqr2()
	}

	fmt.Scanln()
	fmt.Println("Result: ", mas)
}
