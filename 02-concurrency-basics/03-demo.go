package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")

	wg.Add(1) // increment the counter by 1
	go f1(wg) // scheduling f1

	f2()
	wg.Wait() //wait for the counter to become 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
