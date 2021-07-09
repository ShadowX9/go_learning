package main

import "fmt"

func main() {

	counters := map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
		}(counters, i)
	}

	fmt.Scanln()
	fmt.Println("Result: ", counters)
}

// go test -race ./main.go - можно увидеть список горутин, которые обращаются одновременно к одинаковым адресам