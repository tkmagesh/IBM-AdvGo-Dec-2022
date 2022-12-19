package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	for i := 0; i < 10; i++ {
		wg.Add(1) // increment the counter by 1
		go f1()   // scheduling f1
	}
	f2()
	wg.Wait() //wait for the counter to become 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
