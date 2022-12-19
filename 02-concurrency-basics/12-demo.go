package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println("in main")
	fmt.Println(<-ch)
	fmt.Println("in main")
	fmt.Println(<-ch)
	fmt.Println("in main")
	fmt.Println(<-ch)
	fmt.Println("in main")
	fmt.Println(<-ch)
	fmt.Println("in main")
	fmt.Println(<-ch)
}

func fn(ch chan int) {
	time.Sleep(500 * time.Millisecond)
	ch <- 10
	fmt.Println("in fn")
	time.Sleep(500 * time.Millisecond)
	ch <- 20
	fmt.Println("in fn")
	time.Sleep(500 * time.Millisecond)
	ch <- 30
	fmt.Println("in fn")
	time.Sleep(500 * time.Millisecond)
	ch <- 40
	fmt.Println("in fn")
	time.Sleep(500 * time.Millisecond)
	ch <- 50
	fmt.Println("in fn")
}
