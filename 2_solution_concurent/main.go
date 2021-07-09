package main

import (
	"fmt"
	"sync"
)

func main() {

	counters := map[int]int{}

	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			mutex.Lock()
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
			mutex.Unlock()
		}(counters, i)
	}

	fmt.Scanln()
	fmt.Println("Result: ", counters)
}
