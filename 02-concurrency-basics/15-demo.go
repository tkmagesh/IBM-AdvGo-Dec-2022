package main

import (
	"fmt"
	"math/rand"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	rand.Seed(time.Now().Unix())
	go fn(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

//producer
func fn(ch chan int) {
	count := rand.Intn(50)
	fmt.Println("count :", count)
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
