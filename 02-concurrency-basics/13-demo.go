package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	count := 5
	go fn(ch, count)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
}

func fn(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
}
