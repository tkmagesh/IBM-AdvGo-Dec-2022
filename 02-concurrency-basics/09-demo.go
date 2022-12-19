package main

import (
	"fmt"
	"sync"
)

/* Share memory by communicating (using channels) */
func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
