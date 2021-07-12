package main

import (
	"reflect"
	"testing"
)

func TestExecutePipeline(t *testing.T) {
	var gen = func(in, out chan interface{}) {
		for i := 0; i < 10; i++ {
			out <- i
		}
	}
	var takeOdd = func(in, out chan interface{}) {
		var count = 0
		for v := range in {
			if count%2 != 0 {
				out <- v
			}
			count++
		}
	}
	var square = func(in, out chan interface{}) {
		for v := range in {
			var val = v.(int)
			out <- val * val
		}
	}

	var res = make([]int, 0, 10)
	var collect = func(in, out chan interface{}) {
		for v := range in {
			res = append(res, v.(int))
		}
	}

	ExecutePipeline(
		job(gen),
		job(takeOdd),
		job(square),
		job(collect),
	)

	var expected = []int{1, 9, 25, 49, 81}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("wrong answer. Expected: %v Got: %v", expected, res)
	}
}
