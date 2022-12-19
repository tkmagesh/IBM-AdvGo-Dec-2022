package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}
	rand.Seed(7)
	fmt.Println("main started")
	noOfGoroutines := flag.Int("count", 0, "Number of goroutines to run")
	noOfCpu := flag.Int("cpu", runtime.NumCPU(), "Number of CPUs to use")
	flag.Parse()
	runtime.GOMAXPROCS(*noOfCpu)
	fmt.Printf("Starting %d goroutines using %d CPUs\n", *noOfGoroutines, *noOfCpu)
	fmt.Println("Hit ENTER to start")
	fmt.Scanln()

	for i := 1; i <= *noOfGoroutines; i++ {
		wg.Add(1)    // increment the counter by 1
		go f1(i, wg) // scheduling f1
	}

	f2()
	wg.Wait() //wait for the counter to become 0
	fmt.Println("main completed")
	fmt.Println("Hit ENTER to exit")
	fmt.Scanln()
	elapsed := time.Since(start)
	fmt.Println("Time taken :", elapsed)
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	/*
		result := 0
		for i := 0; i < 10000; i++ {
			for j := 0; j < 10000; j++ {
				for k := 0; k < 100; k++ {
					result += 1
				}
			}
		}
	*/
	fmt.Printf("f1[%d] completed\n", id)
}

func f2() {
	fmt.Println("f2 invoked")
}
